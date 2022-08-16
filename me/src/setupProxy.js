const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
    app.use(
        '/v2',
        createProxyMiddleware({
            target: 'http://localhost:9000',
            changeOrigin: true,
        })
    );
    app.use(
        '/file',
        createProxyMiddleware({
            target: 'http://localhost:8282',
            changeOrigin: true,
        })
    );
};