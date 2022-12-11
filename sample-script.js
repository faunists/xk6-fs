import file from 'k6/x/fs';

const root = 'samples';

export default function () {
    const jsFiles = file.walkMatch(root, '*.js');
    console.log(jsFiles)
}