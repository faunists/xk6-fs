build-local:
	xk6 build --with github.com/faunists/xk6-fs@latest=.

run-sample:
	./k6 run sample-script.js