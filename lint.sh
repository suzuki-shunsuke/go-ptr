set -e

echo "+ gofmt"
! git ls-files | grep .go | xargs gofmt -s -d | grep '^'
echo "+ go vet"
go vet $(go list ./... | grep -v /vendor/)
echo "+ golint"
golint ./...
echo "+ gosimple"
! gosimple ./... | grep '^'
echo "+ unused"
unused ./...
echo "+ staticcheck (failure is ignored)"
staticcheck ./... || echo "staticcheck failure is ignored"
