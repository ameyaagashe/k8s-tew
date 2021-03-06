package utils

// Versions
const VERSION_K8S = "1.11.3"
const VERSION_CONFIG = "2.1.0"
const VERSION_ETCD = "3.3.9"
const VERSION_CONTAINERD = "1.1.3"
const VERSION_RUNC = "1.0.0-rc5"
const VERSION_CRICTL = "1.11.1"
const VERSION_GOBETWEEN = "0.6.0"
const VERSION_HELM = "2.9.1"
const VERSION_ARK = "gcr.io/heptio-images/ark:v0.9.6"
const VERSION_MINIO_SERVER = "docker.io/minio/minio:RELEASE.2018-08-18T03-49-57Z"
const VERSION_MINIO_CLIENT = "docker.io/minio/mc:RELEASE.2018-08-18T02-13-04Z"
const VERSION_PAUSE = "k8s.gcr.io/pause:3.1"
const VERSION_COREDNS = "docker.io/coredns/coredns:1.2.0"
const VERSION_ELASTICSEARCH = "docker.io/upmcenterprises/docker-elasticsearch-kubernetes:6.1.3_0"
const VERSION_ELASTICSEARCH_CRON = "docker.io/upmcenterprises/elasticsearch-cron:0.0.3"
const VERSION_ELASTICSEARCH_OPERATOR = "docker.io/upmcenterprises/elasticsearch-operator:0.0.12"
const VERSION_KIBANA = "docker.elastic.co/kibana/kibana-oss:6.1.3"
const VERSION_CEREBRO = "docker.io/upmcenterprises/cerebro:0.6.8"
const VERSION_FLUENT_BIT = "docker.io/fluent/fluent-bit:0.13.0"
const VERSION_CALICO_TYPHA = "quay.io/calico/typha:v0.7.4"
const VERSION_CALICO_NODE = "quay.io/calico/node:v3.1.3"
const VERSION_CALICO_CNI = "quay.io/calico/cni:v3.1.3"
const VERSION_RBD_PROVISIONER = "quay.io/external_storage/rbd-provisioner:v2.1.1-k8s1.11"
const VERSION_CEPH = "docker.io/ceph/daemon:v3.0.7-stable-3.0-mimic-centos-7-x86_64"
const VERSION_HEAPSTER = "k8s.gcr.io/heapster:v1.3.0"
const VERSION_ADDON_RESIZER = "k8s.gcr.io/addon-resizer:1.7"
const VERSION_KUBERNETES_DASHBOARD = "k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.0"
const VERSION_CERT_MANAGER_CONTROLLER = "quay.io/jetstack/cert-manager-controller:v0.4.1"
const VERSION_NGINX_INGRESS_DEFAULT_BACKEND = "k8s.gcr.io/defaultbackend:1.4"
const VERSION_NGINX_INGRESS_CONTROLLER = "quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.18.0"
const VERSION_METRICS_SERVER = "gcr.io/google_containers/metrics-server-amd64:v0.2.1"
const VERSION_PROMETHEUS_OPERATOR = "quay.io/coreos/prometheus-operator:v0.20.0"
const VERSION_PROMETHEUS_CONFIG_RELOADER = "quay.io/coreos/prometheus-config-reloader:v0.20.0"
const VERSION_CONFIGMAP_RELOAD = "quay.io/coreos/configmap-reload:v0.0.1"
const VERSION_KUBE_STATE_METRICS = "gcr.io/google_containers/kube-state-metrics:v1.2.0"
const VERSION_GRAFANA = "docker.io/grafana/grafana:5.0.0"
const VERSION_GRAFANA_WATCHER = "quay.io/coreos/grafana-watcher:v0.0.8"
const VERSION_PROMETHEUS = "quay.io/prometheus/prometheus:v2.2.1"
const VERSION_PROMETHEUS_NODE_EXPORTER = "quay.io/prometheus/node-exporter:v0.15.2"
const VERSION_PROMETHEUS_ALERT_MANAGER = "quay.io/prometheus/alertmanager:v0.15.1"
const VERSION_CSI_ATTACHER = "quay.io/k8scsi/csi-attacher:v0.3.0"
const VERSION_CSI_PROVISIONER = "quay.io/k8scsi/csi-provisioner:v0.3.0"
const VERSION_CSI_DRIVER_REGISTRAR = "quay.io/k8scsi/driver-registrar:v0.3.0"
const VERSION_CSI_CEPH_RBD_PLUGIN = "quay.io/cephcsi/rbdplugin:v0.3.0"
const VERSION_CSI_CEPH_FS_PLUGIN = "quay.io/cephcsi/cephfsplugin:v0.3.0"
const VERSION_MYSQL = "docker.io/library/mysql:5.6"
const VERSION_WORDPRESS = "docker.io/library/wordpress:4.8-apache"

// Settings
const PROJECT_TITLE = "Kubernetes - The Easier Way"
const CLUSTER_NAME = "k8s-tew"
const RSA_SIZE = 2048
const CA_VALIDITY_PERIOD = 20
const CLIENT_VALIDITY_PERIOD = 15
const BASE_DIRECTORY = "assets"
const CLUSTER_DOMAIN = "cluster.local"
const CLUSTER_IP_RANGE = "10.32.0.0/24"
const CALICO_TYPHA_IP = "10.32.0.5"
const CLUSTER_DNS_IP = "10.32.0.10"
const CLUSTER_CIDR = "10.200.0.0/16"
const RESOLV_CONF = "/etc/resolv.conf"
const PUBLIC_NETWORK = "192.168.100.0/24"
const HELM_SERVICE_ACCOUNT = "tiller"
const EMAIL = "k8s-tew@gmail.com"
const DEPLOYMENT_DIRECTORY = "/"
const INGRESS_DOMAIN = "k8s-tew.net"
const INGRESS_SUBDOMAIN_WORDPRESS = "wordpress"

// Ports
const PORT_VIP_RAFT_CONTROLLER uint16 = 16277
const PORT_VIP_RAFT_WORKER uint16 = 16728
const PORT_LOAD_BALANCER uint16 = 16443
const PORT_KUBERNETES_DASHBOARD uint16 = 32443
const PORT_API_SERVER uint16 = 6443
const PORT_CEPH_MANAGER uint16 = 30700
const PORT_CEPH_RADOS_GATEWAY uint16 = 30750
const PORT_MINIO uint16 = 30800
const PORT_GRAFANA uint16 = 30900
const PORT_KIBANA uint16 = 30980
const PORT_CEREBRO uint16 = 30990
const PORT_WORDPRESS uint16 = 30100

// URLs

const K8S_BASE_NAME = "kubernetes-server-linux-amd64"
const K8S_DOWNLOAD_URL = "https://storage.googleapis.com/kubernetes-release/release/v{{.Versions.K8S}}/{{.Filename}}.tar.gz"
const ETCD_BASE_NAME = "etcd-v{{.Versions.Etcd}}-linux-amd64"
const ETCD_DOWNLOAD_URL = "https://github.com/coreos/etcd/releases/download/v{{.Versions.Etcd}}/{{.Filename}}.tar.gz"
const FLANNELD_DOWNLOAD_URL = "https://github.com/coreos/flannel/releases/download/v{{.Versions.Flanneld}}/flanneld-amd64"
const CNI_BASE_NAME = "cni-plugins-amd64-v{{.Versions.CNI}}"
const CNI_DOWNLOAD_URL = "https://github.com/containernetworking/plugins/releases/download/v{{.Versions.CNI}}/{{.Filename}}.tgz"
const CONTAINERD_BASE_NAME = "containerd-{{.Versions.Containerd}}.linux-amd64"
const CONTAINERD_DOWNLOAD_URL = "https://github.com/containerd/containerd/releases/download/v{{.Versions.Containerd}}/{{.Filename}}.tar.gz"
const RUNC_DOWNLOAD_URL = "https://github.com/opencontainers/runc/releases/download/v{{.Versions.Runc}}/runc.amd64"
const CRICTL_BASE_NAME = "crictl-v{{.Versions.CriCtl}}-linux-amd64"
const CRICTL_DOWNLOAD_URL = "https://github.com/kubernetes-incubator/cri-tools/releases/download/v{{.Versions.CriCtl}}/{{.Filename}}.tar.gz"
const GOBETWEEN_BASE_NAME = "gobetween_{{.Versions.Gobetween}}_linux_amd64"
const GOBETWEEN_DOWNLOAD_URL = "https://github.com/yyyar/gobetween/releases/download/{{.Versions.Gobetween}}/{{.Filename}}.tar.gz"
const HELM_BASE_NAME = "helm-v{{.Versions.Helm}}-linux-amd64"
const HELM_DOWNLOAD_URL = "https://storage.googleapis.com/kubernetes-helm/{{.Filename}}.tar.gz"
const ARK_BASE_NAME = "ark-{{.Versions.Ark | image_tag}}-linux-amd64"
const ARK_DOWNLOAD_URL = "https://github.com/heptio/ark/releases/download/{{.Versions.Ark | image_tag}}/{{.Filename}}.tar.gz"

// Config
const CONFIG_FILENAME = "config.yaml"

// Node Labels
const NODE_BOOTSTRAPPER = "bootstrapper"
const NODE_CONTROLLER = "controller"
const NODE_WORKER = "worker"
const NODE_STORAGE = "storage"

// Features
const FEATURE_STORAGE = "storage"
const FEATURE_MONITORING = "monitoring"
const FEATURE_LOGGING = "logging"
const FEATURE_BACKUP = "backup"
const FEATURE_SHOWCASE = "showcase"
const FEATURE_INGRESS = "ingress"
const FEATURE_PACKAGING = "packaging"

// OS
const OS_UBUNTU = "ubuntu"
const OS_UBUNTU_18_04 = "ubuntu/18.04"
const OS_CENTOS = "centos"
const OS_CENTOS_7_5 = "centos/7.5"

// Sub-Directories
const TEMPORARY_SUBDIRECTORY = "tmp"
const CONFIG_SUBDIRECTORY = "etc"
const SYSTEMD_SUBDIRECTORY = "systemd"
const SYSTEM_SUBDIRECTORY = "system"
const K8S_TEW_SUBDIRECTORY = "k8s-tew"
const CERTIFICATES_SUBDIRECTORY = "ssl"
const OPTIONAL_SUBDIRECTORY = "opt"
const VARIABLE_SUBDIRECTORY = "var"
const LOGGING_SUBDIRECTORY = "log"
const LIBRARY_SUBDIRECTORY = "lib"
const RUN_SUBDIRECTORY = "run"
const BINARY_SUBDIRECTORY = "bin"
const K8S_SUBDIRECTORY = "k8s"
const ETCD_SUBDIRECTORY = "etcd"
const CRI_SUBDIRECTORY = "cri"
const CNI_SUBDIRECTORY = "cni"
const KUBECONFIG_SUBDIRECTORY = "kubeconfig"
const SECURITY_SUBDIRECTORY = "security"
const SETUP_SUBDIRECTORY = "setup"
const CONTAINERD_SUBDIRECTORY = "containerd"
const PROFILE_D_SUBDIRECTORY = "profile.d"
const LOAD_BALANCER_SUBDIRECTORY = "lb"
const HELM_SUBDIRECTORY = "helm"
const KUBELET_SUBDIRECTORY = "kubelet"
const PODS_SUBDIRECTORY = "pods"
const MANIFESTS_SUBDIRECTORY = "manifests"
const CEPH_SUBDIRECTORY = "ceph"
const CEPH_BOOTSTRAP_MDS_SUBDIRECTORY = "bootstrap-mds"
const CEPH_BOOTSTRAP_OSD_SUBDIRECTORY = "bootstrap-osd"
const CEPH_BOOTSTRAP_RBD_SUBDIRECTORY = "bootstrap-rbd"
const CEPH_BOOTSTRAP_RGW_SUBDIRECTORY = "bootstrap-rgw"
const ARK_SUBDIRECTORY = "ark"
const BASH_COMPLETION_SUBDIRECTORY = "bash_completion.d"
const HOST_SUBDIRECTORY = "host"
const PLUGINS_SUBDIRECTORY = "plugins"
const CSI_CEPHFS_PLUGIN = "csi-cephfsplugin"
const CSI_RBD_PLUGIN = "csi-rbdplugin"

// Directories
const CONFIG_DIRECTORY = "config"
const CERTIFICATES_DIRECTORY = "certificates"
const CNI_CONFIG_DIRECTORY = "cni-config"
const CRI_CONFIG_DIRECTORY = "cri-config"
const K8S_SECURITY_CONFIG_DIRECTORY = "security-config"
const K8S_CONFIG_DIRECTORY = "k8s-config"
const K8S_KUBE_CONFIG_DIRECTORY = "kube-config"
const K8S_SETUP_CONFIG_DIRECTORY = "setup-config"
const BINARIES_DIRECTORY = "binaries"
const K8S_BINARIES_DIRECTORY = "k8s-binaries"
const ETCD_BINARIES_DIRECTORY = "etcd-binaries"
const CNI_BINARIES_DIRECTORY = "cni-binaries"
const CRI_BINARIES_DIRECTORY = "cri-binaries"
const DYNAMIC_DATA_DIRECTORY = "dynamic-data"
const ETCD_DATA_DIRECTORY = "etcd-data"
const CONTAINERD_DATA_DIRECTORY = "containerd-data"
const LOGGING_DIRECTORY = "logging"
const SERVICE_DIRECTORY = "service"
const CONTAINERD_STATE_DIRECTORY = "containerd-state"
const ABSOLUTE_CONTAINERD_STATE_DIRECTORY = "absolute-containerd-state"
const PROFILE_DIRECTORY = "profile"
const GOBETWEEN_BINARIES_DIRECTORY = "gobetween-binaries"
const GOBETWEEN_CONFIG_DIRECTORY = "gobetween-config"
const HELM_DATA_DIRECTORY = "helm-data"
const KUBELET_DATA_DIRECTORY = "kubelet-data"
const PODS_DATA_DIRECTORY = "pods-data"
const TEMPORARY_DIRECTORY = "temporary"
const K8S_MANIFESTS_DIRECTORY = "kubelet-manifests"
const CEPH_DIRECTORY = "ceph"
const CEPH_CONFIG_DIRECTORY = "ceph-config"
const CEPH_DATA_DIRECTORY = "ceph-data"
const CEPH_BOOTSTRAP_MDS_DIRECTORY = "bootstrap-mds"
const CEPH_BOOTSTRAP_OSD_DIRECTORY = "bootstrap-osd"
const CEPH_BOOTSTRAP_RBD_DIRECTORY = "bootstrap-rbd"
const CEPH_BOOTSTRAP_RGW_DIRECTORY = "bootstrap-rgw"
const ARK_BINARIES_DIRECTORY = "ark"
const BASH_COMPLETION_DIRECTORY = "bash-completion"
const HOST_BINARIES_DIRECTORY = "host-binaries"
const CEPH_RBD_PLUGIN_DIRECTORY = "ceph-rbd-plugin"
const CEPH_FS_PLUGIN_DIRECTORY = "ceph-fs-plugin"
const KUBELET_PLUGINS_DIRECTORY = "kubelet-plugins"

// Binaries
const K8S_TEW_BINARY = "k8s-tew"

// Helm Binary
const HELM_BINARY = "helm"

// ContainerD Binaries
const CONTAINERD_BINARY = "containerd"
const CONTAINERD_SHIM_BINARY = "containerd-shim"
const CTR_BINARY = "ctr"
const RUNC_BINARY = "runc"
const CRICTL_BINARY = "crictl"

// Etcd Binaries
const ETCD_BINARY = "etcd"
const ETCDCTL_BINARY = "etcdctl"

// K8S Binaries
const KUBECTL_BINARY = "kubectl"
const KUBE_APISERVER_BINARY = "kube-apiserver"
const KUBE_CONTROLLER_MANAGER_BINARY = "kube-controller-manager"
const KUBELET_BINARY = "kubelet"
const KUBE_PROXY_BINARY = "kube-proxy"
const KUBE_SCHEDULER_BINARY = "kube-scheduler"

// Gobeween Binary
const GOBETWEEN_BINARY = "gobetween"

// Ark Binaries
const ARK_BINARY = "ark"

// Certificates
const CA_PEM = "ca.pem"
const CA_KEY_PEM = "ca-key.pem"
const KUBERNETES_PEM = "kubernetes.pem"
const KUBERNETES_KEY_PEM = "kubernetes-key.pem"
const ADMIN_PEM = "admin.pem"
const ADMIN_KEY_PEM = "admin-key.pem"
const PROXY_PEM = "proxy.pem"
const PROXY_KEY_PEM = "proxy-key.pem"
const CONTROLLER_MANAGER_PEM = "controller-manager.pem"
const CONTROLLER_MANAGER_KEY_PEM = "controller-manager-key.pem"
const SCHEDULER_PEM = "scheduler.pem"
const SCHEDULER_KEY_PEM = "scheduler-key.pem"
const KUBELET_PEM = "kubelet-{{.Name}}.pem"
const KUBELET_KEY_PEM = "kubelet-{{.Name}}-key.pem"
const SERVICE_ACCOUNT_PEM = "service-account.pem"
const SERVICE_ACCOUNT_KEY_PEM = "service-account-key.pem"
const FLANNELD_PEM = "flanneld.pem"
const FLANNELD_KEY_PEM = "flanneld-key.pem"
const VIRTUAL_IP_PEM = "virtual-ip.pem"
const VIRTUAL_IP_KEY_PEM = "virtual-ip-key.pem"
const AGGREGATOR_PEM = "aggregator.pem"
const AGGREGATOR_KEY_PEM = "aggregator-key.pem"

// Kubeconfig
const ADMIN_KUBECONFIG = "admin.kubeconfig"
const CONTROLLER_MANAGER_KUBECONFIG = "controller-manager.kubeconfig"
const SCHEDULER_KUBECONFIG = "scheduler.kubeconfig"
const PROXY_KUBECONFIG = "proxy.kubeconfig"
const KUBELET_KUBECONFIG = "kubelet-{{.Name}}.kubeconfig"

// Security
const ENCRYPTION_CONFIG = "encryption-config.yaml"

// Containerd
const CONTAINERD_CONFIG = "config-{{.Name}}.toml"
const CONTAINERD_SOCK = "containerd.sock"

// K8S Config
const K8S_KUBELET_SETUP = "kubelet-setup.yaml"
const K8S_ADMIN_USER_SETUP = "admin-user-setup.yaml"
const K8S_HELM_USER_SETUP = "helm-user-setup.yaml"
const K8S_KUBE_SCHEDULER_CONFIG = "kube-scheduler-config.yaml"
const K8S_KUBELET_CONFIG = "kubelet-{{.Name}}-config.yaml"
const K8S_COREDNS_SETUP = "coredns-setup.yaml"
const K8S_CALICO_SETUP = "calico-setup.yaml"
const K8S_ELASTICSEARCH_OPERATOR_SETUP = "elasticsearch-operator-setup.yaml"
const K8S_EFK_SETUP = "efk-setup.yaml"
const K8S_ARK_SETUP = "ark-setup.yaml"
const K8S_HEAPSTER_SETUP = "heapster-setup.yaml"
const K8S_KUBERNETES_DASHBOARD_SETUP = "kubernetes-dashboard-setup.yaml"
const K8S_CERT_MANAGER_SETUP = "cert-manager-setup.yaml"
const K8S_NGINX_INGRESS_SETUP = "nginx-ingress-setup.yaml"
const K8S_METRICS_SERVER_SETUP = "metrics-server-setup.yaml"
const K8S_PROMETHEUS_OPERATOR_SETUP = "prometheus-operator-setup.yaml"
const K8S_KUBE_PROMETHEUS_SETUP = "kube-prometheus-setup.yaml"
const K8S_KUBE_PROMETHEUS_DATASOURCE_SETUP = "kube-prometheus-datasource-setup.yaml"
const K8S_KUBE_PROMETHEUS_KUBERNETES_CLUSTER_STATUS_DASHBOARD_SETUP = "kube-prometheus-kubernetes-cluster-status-dashboard-setup.yaml"
const K8S_KUBE_PROMETHEUS_PODS_DASHBOARD_SETUP = "kube-prometheus-pods-dashboard-setup.yaml"
const K8S_KUBE_PROMETHEUS_DEPLOYMENT_DASHBOARD_SETUP = "kube-prometheus-deployment-dashboard-setup.yaml"
const K8S_KUBE_PROMETHEUS_KUBERNETES_CONTROL_PLANE_STATUS_DASHBOARD_SETUP = "kube-prometheus-kubernetes-control-plane-status-dashboard-setup.yaml"
const K8S_KUBE_PROMETHEUS_STATEFULSET_DASHBOARD_SETUP = "kube-prometheus-statefulset-dashboard-setup.yaml"
const K8S_KUBE_PROMETHEUS_KUBERNETES_CAPACITY_PLANNING_DASHBOARD_SETUP = "kube-prometheus-kubernetes-capacity-planning-dashboard-setup.yaml"
const K8S_KUBE_PROMETHEUS_KUBERNETES_RESOURCE_REQUESTS_DASHBOARD_SETUP = "kube-prometheus-kubernetes-resource-requests-dashboard-setup.yaml"
const K8S_KUBE_PROMETHEUS_KUBERNETES_CLUSTER_HEALTH_DASHBOARD_SETUP = "kube-prometheus-kubernetes-cluster-health-dashboard-setup.yaml"
const K8S_KUBE_PROMETHEUS_NODES_DASHBOARD_SETUP = "kube-prometheus-nodes-dashboard-setup.yaml"
const WORDPRESS_SETUP = "wordpress-setup.yaml"

// Gobetween Config
const GOBETWEEN_CONFIG = "config.toml"

// Profile
const K8S_TEW_PROFILE = "k8s-tew.sh"

// Bash Completion
const BASH_COMPLETION_K8S_TEW = "k8s-tew.bash-completion"
const BASH_COMPLETION_KUBECTL = "kubectl.bash-completion"
const BASH_COMPLETION_HELM = "helm.bash-completion"
const BASH_COMPLETION_ARK = "ark.bash-completion"
const BASH_COMPLETION_CRICTL = "crictl.bash-completion"

// Logging
const AUDIT_LOG = "audit.log"

// Deployment
const DEPLOYMENT_USER = "root"

// Service
const SERVICE_NAME = "k8s-tew"
const SERVICE_CONFIG = SERVICE_NAME + ".service"

// Ceph
const CEPH_RBD_POOL_NAME = "cephrbd"
const CEPH_FS_POOL_NAME = "cephfs"
const CEPH_CONFIG = "ceph.conf"
const CEPH_CLIENT_ADMIN_KEYRING = "ceph.client.admin.keyring"
const CEPH_MONITOR_KEYRING = "ceph.mon.keyring"
const CEPH_KEYRING = "ceph.keyring"
const CEPH_BOOTSTRAP_MDS_KEYRING = "ceph.bootstrap.mds.keyring"
const CEPH_BOOTSTRAP_OSD_KEYRING = "ceph.bootstrap.osd.keyring"
const CEPH_BOOTSTRAP_RBD_KEYRING = "ceph.bootstrap.rbd.keyring"
const CEPH_BOOTSTRAP_RGW_KEYRING = "ceph.bootstrap.rgw.keyring"
const CEPH_SECRETS = "ceph-secrets.yaml"
const CEPH_SETUP = "ceph-setup.yaml"
const CEPH_CSI = "ceph-csi.yaml"

// Cluster Issuer
const LETSENCRYPT_CLUSTER_ISSUER = "letsencrypt-cluster-issuer.yaml"

// Environment variables
const K8S_TEW_BASE_DIRECTORY = "K8S_TEW_BASE_DIRECTORY"

// Virtual IP Manager
const ELECTION_NAMESPACE = "/k8s-tew"
const ELECTION_CONTROLLER = "/controller-vip-manager"
const ELECTION_WORKER = "/worker-vip-manager"

// Common Names
const CN_ADMIN = "admin"
const CN_AGGREGATOR = "aggregator"
const CN_SYSTEM_KUBE_CONTROLLER_MANAGER = "system:kube-controller-manager"
const CN_SYSTEM_KUBE_SCHEDULER = "system:kube-scheduler"
const CN_SYSTEM_KUBE_PROXY = "system:kube-proxy"
const CN_SYSTEM_NODE_PREFIX = "system:node:%s"

// Templates
const TEMPLATE_CONTAINERD_TOML = "k8s/cri/containerd.toml"
const TEMPLATE_K8S_TEW_SERVICE = "system/k8s-tew.service"
const TEMPLATE_K8S_TEW_PROFILE = "system/k8s-tew.sh"
const TEMPLATE_ENVIRONMENT = "system/environment.sh"
const TEMPLATE_GOBETWEEN_TOML = "k8s/lb/gobetween.toml"
const TEMPLATE_KUBE_SCHEDULER_CONFIGURATION = "k8s/kube-scheduler-configuration.yaml"
const TEMPLATE_KUBELET_CONFIGURATION = "k8s/kubelet-configuration.yaml"
const TEMPLATE_ENCRYPTION_CONFIG = "k8s/encryption-config.yaml"
const TEMPLATE_KUBECONFIG = "k8s/kubeconfig.yaml"
const TEMPLATE_SERVICE_ACCOUNT = "k8s/service-account.yaml"
const TEMPLATE_KUBELET_SETUP = "k8s/setup/kubelet-setup.yaml"
const TEMPLATE_CEPH_CLIENT_KEYRING = "ceph/client.keyring"
const TEMPLATE_CEPH_CLIENT_ADMIN_KEYRING = "ceph/client-admin.keyring"
const TEMPLATE_CEPH_MONITOR_KEYRING = "ceph/monitor.keyring"
const TEMPLATE_CEPH_CONFIG = "ceph/ceph.conf"
const TEMPLATE_CEPH_SECRETS = "k8s/setup/storage/ceph-secrets.yaml"
const TEMPLATE_CEPH_SETUP = "k8s/setup/storage/ceph-setup.yaml"
const TEMPLATE_CEPH_CSI = "k8s/setup/storage/ceph-csi.yaml"
const TEMPLATE_LETSENCRYPT_CLUSTER_ISSUER_SETUP = "k8s/setup/ingress/letsencrypt-cluster-issuer.yaml"
const TEMPLATE_COREDNS_SETUP = "k8s/setup/dns/coredns.yaml"
const TEMPLATE_CALICO_SETUP = "k8s/setup/networking/calico.yaml"
const TEMPLATE_ELASTICSEARCH_OPERATOR_SETUP = "k8s/setup/logging/elasticsearch-operator.yaml"
const TEMPLATE_EFK_SETUP = "k8s/setup/logging/efk.yaml"
const TEMPLATE_ARK_SETUP = "k8s/setup/backup/ark.yaml"
const TEMPLATE_HEAPSTER_SETUP = "k8s/setup/monitoring/heapster.yaml"
const TEMPLATE_KUBERNETES_DASHBOARD_SETUP = "k8s/setup/management/kubernetes-dashboard.yaml"
const TEMPLATE_CERT_MANAGER_SETUP = "k8s/setup/networking/cert-manager.yaml"
const TEMPLATE_NGINX_INGRESS_SETUP = "k8s/setup/networking/nginx-ingress.yaml"
const TEMPLATE_METRICS_SERVER_SETUP = "k8s/setup/monitoring/metrics-server.yaml"
const TEMPLATE_PROMETHEUS_OPERATOR_SETUP = "k8s/setup/monitoring/prometheus-operator.yaml"
const TEMPLATE_KUBE_PROMETHEUS_SETUP = "k8s/setup/monitoring/kube-prometheus.yaml"
const TEMPLATE_KUBE_PROMETHEUS_DATASOURCE_SETUP = "k8s/setup/monitoring/kube-prometheus-datasource.yaml"
const TEMPLATE_KUBE_PROMETHEUS_KUBERNETES_CLUSTER_STATUS_DASHBOARD_SETUP = "k8s/setup/monitoring/kube-prometheus-kubernetes-cluster-status-dashboard.yaml"
const TEMPLATE_KUBE_PROMETHEUS_PODS_DASHBOARD_SETUP = "k8s/setup/monitoring/kube-prometheus-pods-dashboard.yaml"
const TEMPLATE_KUBE_PROMETHEUS_DEPLOYMENT_DASHBOARD_SETUP = "k8s/setup/monitoring/kube-prometheus-deployment-dashboard.yaml"
const TEMPLATE_KUBE_PROMETHEUS_KUBERNETES_CONTROL_PLANE_STATUS_DASHBOARD_SETUP = "k8s/setup/monitoring/kube-prometheus-kubernetes-control-plane-status-dashboard.yaml"
const TEMPLATE_KUBE_PROMETHEUS_STATEFULSET_DASHBOARD_SETUP = "k8s/setup/monitoring/kube-prometheus-statefulset-dashboard.yaml"
const TEMPLATE_KUBE_PROMETHEUS_KUBERNETES_CAPACITY_PLANNING_DASHBOARD_SETUP = "k8s/setup/monitoring/kube-prometheus-kubernetes-capacity-planning-dashboard.yaml"
const TEMPLATE_KUBE_PROMETHEUS_KUBERNETES_RESOURCE_REQUESTS_DASHBOARD_SETUP = "k8s/setup/monitoring/kube-prometheus-kubernetes-resource-requests-dashboard.yaml"
const TEMPLATE_KUBE_PROMETHEUS_KUBERNETES_CLUSTER_HEALTH_DASHBOARD_SETUP = "k8s/setup/monitoring/kube-prometheus-kubernetes-cluster-health-dashboard.yaml"
const TEMPLATE_KUBE_PROMETHEUS_NODES_DASHBOARD_SETUP = "k8s/setup/monitoring/kube-prometheus-nodes-dashboard.yaml"
const TEMPLATE_WORDPRESS_SETUP = "k8s/setup/miscellaneous/wordpress.yaml"
