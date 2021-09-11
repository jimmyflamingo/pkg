build: 
	tags="$(grep -I  -r '// +build' . | \
                grep -v '^./vendor/' | \
                grep -v '^./hack/' | \
                grep -v '^./third_party' | \
                cut -f3 -d' ' | \
                sort | uniq | \
                grep -v '^!' | \
                tr '\n' ' ')"
	echo "Building with tags: ${tags}"
	go test -vet=off -tags "${tags}" -exec echo ./...