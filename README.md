# Example

## Setup

Pretending to import the external project by using go workspaces.

```bash
mkdir internal external

cd external/
go mod init example.io/other
kubebuilder init --domain example.io --repo example.io/other --plugins=go/v4-alpha
kubebuilder create api --group=other --version v2 --kind External --controller --resource
cd -

cd internal/
go mod init example.io/demo
kubebuilder init --domain example.io --repo example.io/demo --plugins=go/v4-alpha
kubebuilder create api --group=demo --version v1 --kind Internal --controller --resource
kubebuilder create api --group=other --version v2 --kind External --controller --resource=false
cd -

go work init
go work use external/
go work use internal/
```
