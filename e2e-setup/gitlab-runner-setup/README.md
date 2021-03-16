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
