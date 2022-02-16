#!/usr/bin/env bash

package=$1
package_name=$2
if [[ -z "$package" || -z "$package_name" ]]; then
  echo "usage: $0 <package-name> $1 <name>"
  exit 1
fi
	
platforms=("windows/amd64" "windows/386" "darwin/amd64")

mkdir bin
cd bin
package="../$package"

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name'-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done
