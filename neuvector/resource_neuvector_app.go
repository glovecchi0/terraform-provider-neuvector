package neuvector

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationCreate,
		ReadContext:   resourceApplicationRead,
		UpdateContext: resourceApplicationUpdate,
		DeleteContext: resourceApplicationDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "neuvector",
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cattle-neuvector-system",
			},
                        "app_version": {
                                Type:     schema.TypeString,
                                Optional: true,
                        },
			"controller_replicas": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
                        "controller_env": {
                                Type:     schema.TypeList,
                                Optional: true,
                                Elem: &schema.Resource{
                                        Schema: map[string]*schema.Schema{
                                                "name": {
                                                        Type:     schema.TypeString,
                                                        Required: true,
                                                },
                                                "value": {
                                                        Type:     schema.TypeString,
                                                        Required: true,
                                                },
                                        },
                                },
                        },
			"controller_node_selector": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"controller_secret_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"controller_secret_password": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				RequiredWith:  []string{"controller_secret_enabled"},
				ForceNew:      true,
			},
			"manager_svc_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ClusterIP",
			},
			"cve_scanner_replicas": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"cve_scanner_node_selector": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"resources_limits_cpu": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "400m",
			},
			"resources_limits_memory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2792Mi",
			},
			"resources_requests_cpu": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "100m",
			},
			"resources_requests_memory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2280Mi",
			},
			"containerd_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"containerd_path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/run/containerd/containerd.sock",
			},
			"kubeconfig_path": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if err := addHelmRepository(); err != nil {
		return diag.FromErr(err)
	}

	values, err := createValues(d)
	if err != nil {
		return diag.FromErr(err)
	}

	valuesFile, err := createTempValuesFile(values)
	if err != nil {
		return diag.FromErr(err)
	}
	defer os.Remove(valuesFile.Name())

	if err := executeHelmCommand("install", d.Get("name").(string), d.Get("namespace").(string), d.Get("kubeconfig_path").(string), valuesFile.Name()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("neuvector")
	return nil
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implement the read function to refresh resource state
	return nil
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if err := addHelmRepository(); err != nil {
		return diag.FromErr(err)
	}

	values, err := createValues(d)
	if err != nil {
		return diag.FromErr(err)
	}

	valuesFile, err := createTempValuesFile(values)
	if err != nil {
		return diag.FromErr(err)
	}
	defer os.Remove(valuesFile.Name())

	if err := executeHelmCommand("upgrade", d.Get("name").(string), d.Get("namespace").(string), d.Get("kubeconfig_path").(string), valuesFile.Name()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("neuvector")
	return nil
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)
	kubeconfigPath := d.Get("kubeconfig_path").(string)

	if err := executeHelmCommand("uninstall", name, namespace, kubeconfigPath, ""); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func addHelmRepository() error {
	cmd := exec.Command("helm", "repo", "add", "neuvector", "https://neuvector.github.io/neuvector-helm/")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add Helm repository: %w", err)
	}
	cmd = exec.Command("helm", "repo", "update")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to update Helm repositories: %w", err)
	}
	return nil
}

func createValues(d *schema.ResourceData) (string, error) {
	envList := d.Get("controller_env").([]interface{})
	envVars := make([]map[string]string, len(envList))
	for i, env := range envList {
		envMap := env.(map[string]interface{})
		envVars[i] = map[string]string{
			"name":  envMap["name"].(string),
			"value": envMap["value"].(string),
		}
	}

	values := struct {
                NeuVectorVersion        string
		ControllerReplicas      int
                ControllerEnv           []map[string]string
		ControllerNodeSelector  map[string]interface{}
                SecretEnabled           bool
                NeuVectorAdminPwd       string
		ManagerSvcType          string
		CVEScannerReplicas      int
		CVEScannerNodeSelector  map[string]interface{}
		ResourcesLimitsCPU      string
		ResourcesLimitsMemory   string
		ResourcesRequestsCPU    string
		ResourcesRequestsMemory string
		ContainerdEnabled       bool
		ContainerdPath          string
	}{
                NeuVectorVersion:        d.Get("app_version").(string),
		ControllerReplicas:      d.Get("controller_replicas").(int),
                ControllerEnv:           envVars,
		ControllerNodeSelector:  d.Get("controller_node_selector").(map[string]interface{}),
                SecretEnabled:           d.Get("controller_secret_enabled").(bool),
                NeuVectorAdminPwd:       d.Get("controller_secret_password").(string),
		ManagerSvcType:          d.Get("manager_svc_type").(string),
		CVEScannerReplicas:      d.Get("cve_scanner_replicas").(int),
		CVEScannerNodeSelector:  d.Get("cve_scanner_node_selector").(map[string]interface{}),
		ResourcesLimitsCPU:      d.Get("resources_limits_cpu").(string),
		ResourcesLimitsMemory:   d.Get("resources_limits_memory").(string),
		ResourcesRequestsCPU:    d.Get("resources_requests_cpu").(string),
		ResourcesRequestsMemory: d.Get("resources_requests_memory").(string),
		ContainerdEnabled:       d.Get("containerd_enabled").(bool),
		ContainerdPath:          d.Get("containerd_path").(string),
	}

	tmpl, err := template.New("values").Parse(valuesTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse Helm values template: %s", err)
	}

	var valuesBuffer bytes.Buffer
	if err := tmpl.Execute(&valuesBuffer, values); err != nil {
		return "", fmt.Errorf("failed to execute Helm values template: %s", err)
	}

	return valuesBuffer.String(), nil
}

func createTempValuesFile(values string) (*os.File, error) {
	valuesFile, err := os.CreateTemp("", "values-*.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary values file: %s", err)
	}

	if _, err := valuesFile.WriteString(values); err != nil {
		return nil, fmt.Errorf("failed to write to temporary values file: %s", err)
	}

	return valuesFile, nil
}

func executeHelmCommand(action, name, namespace, kubeconfigPath, valuesFilePath string) error {
	cmdArgs := []string{
		action, name,
		"--namespace", namespace,
		"--kubeconfig", kubeconfigPath,
	}

	if action != "uninstall" {
		cmdArgs = append(cmdArgs, "neuvector/core")
	}

	if valuesFilePath != "" {
		cmdArgs = append(cmdArgs, "-f", valuesFilePath)
	}

	cmd := exec.Command("helm", cmdArgs...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute Helm command: %s\nOutput: %s", err, string(output))
	}

	return nil
}

const valuesTemplate = `
{{- if .NeuVectorVersion }}
tag: {{.NeuVectorVersion}}
{{- end }}
controller:
  replicas: {{.ControllerReplicas}}
  env:
  {{- if .ControllerEnv }}
  {{- range .ControllerEnv }}
    - name: {{ .name }}
      value: {{ .value }}
  {{- end }}
  {{- end }}
  nodeSelector:
  {{- range $key, $value := .ControllerNodeSelector }}
    {{ $key }}: {{ $value }}
  {{- end }}
  secret:
    enabled: {{.SecretEnabled}}
    data:
      userinitcfg.yaml: 
        always_reload: true
        users:
        -
          Fullname: admin
          Password: {{.NeuVectorAdminPwd}}
          Role: admin
manager:
  svc:
    type: {{.ManagerSvcType}}
cve:
  scanner:
    replicas: {{.CVEScannerReplicas}}
    nodeSelector:
    {{- range $key, $value := .CVEScannerNodeSelector }}
      {{ $key }}: {{ $value }}
    {{- end }}
resources:
  limits:
    cpu: {{.ResourcesLimitsCPU}}
    memory: {{.ResourcesLimitsMemory}}
  requests:
    cpu: {{.ResourcesRequestsCPU}}
    memory: {{.ResourcesRequestsMemory}}
containerd:
  enabled: {{.ContainerdEnabled}}
  path: {{.ContainerdPath}}
`
