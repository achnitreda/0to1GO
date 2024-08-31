#!/usr/bin/env bash 

# this is a better approach 
# -> env command to find bash by searching the user's PATH.
# how do u know that bash is in /bin (many linux distros uses it)
# But this approach is the better way. 

IMAGE="stat-bin"

# -> linear-stats|math-skills
VALID_BIN=$(ls -xm ./bin | sed 's:, :|:') 



help() {
	cat <<EOF
Run porgram as follow 
./run.sh [$VALID_BIN]
EOF
}

exec_stat_bin() {
	if echo "$VALID_BIN" | grep -q "$1"; then
		docker run -v "$(pwd):$(pwd)" -w "$(pwd)" --rm "$IMAGE" "/stat-bin/$1"
	else
		echo "Error: binary not found!"
		help
		exit 1
	fi
}

if [[ $# -ne 1 ]]; then
	help
	exit 1
fi

if ! docker images | grep -q "$IMAGE"; then
	docker build -t "$IMAGE" .
fi

exec_stat_bin "$1"
