# xk6-fs
[k6](https://github.com/grafana/k6) extension for dealing with the file system, implemented using the
[xk6](https://github.com/grafana/xk6) system.

## Build
```shell
xk6 build --with github.com/faunists/xk6-fs@latest
```

## Example
```javascript
import file from 'k6/x/fs';

const root = '.';

export default function () {
    file.walkMatch(root, '*.js');
}
```

## Run sample script
```shell
make run-sample
```
