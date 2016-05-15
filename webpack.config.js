module.exports = {
    entry: ['./app/main.ts'],
    output: {
        path: './assets/js',
        filename: 'client.bundle.js',
        pathinfo: true,
        sourcePrefix: ''
    },
    resolve: {
        extensions: ['', '.ts', '.tsx', '.js']
    },
    module: {
        loaders: [
            {
                test: /\.tsx?$/,
                loader: 'ts-loader',
                exclude: /node_modules/
            }
        ]
    },
    ts: {
        configFileName: './tsconfig.json'
    }
};