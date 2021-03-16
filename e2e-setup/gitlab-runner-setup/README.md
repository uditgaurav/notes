## Setup GitLab Runner

Run the following commands inside the target VM.

1. Install docker

```bash
sudo su -
apt-get update && apt-get install -y docker.io
```

2. Installing the latest GitLab Runner

```bash
curl -L https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh | sudo bash
export GITLAB_RUNNER_DISABLE_SKEL=true; sudo -E apt-get install gitlab-runner
```
3. Register the Runner


```bash
gitlab-runner register
systemctl restart gitlab-runner
```

4. Give all permission to gitlab-runner user

```bash
visudo
gitlab-runner ALL=(ALL) NOPASSWD: ALL
```

- Checkout the `config.toml` file which is configuration file for GitLab runner

```bash
sudo su - gitlab-runner
sudo vi /etc/gitlab-runner/config.toml
```

Configuration file:
```yaml
concurrent = 1
check_interval = 0

[session_server]
  session_timeout = 1800

[[runners]]
  name = "kubernetes runner for component test"
  url = "https://gitlab.mayadata.io/"
  token = "XXXXXXXXXX"
  executor = "kubernetes"
  [runners.custom_build_dir]
  [runners.cache]
    [runners.cache.s3]
    [runners.cache.gcs]
    [runners.cache.azure]
  [runners.kubernetes]
    host = ""
    bearer_token_overwrite_allowed = false
    image = "litmuschaos/litmus-e2e:ci"
    namespace = "default"
    pull_policy = "always"
    namespace_overwrite_allowed = ""
    privileged = false
    service_account = "litmus-runner"
    service_account_overwrite_allowed = ""
    pod_annotations_overwrite_allowed = ""
    [runners.kubernetes.affinity]
    [runners.kubernetes.pod_security_context]
    [runners.kubernetes.volumes]
 ```

