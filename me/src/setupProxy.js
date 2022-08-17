const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
    app.use(
        '/v2',
        createProxyMiddleware({
            target: 'https://me-server-7t4azswppa-du.a.run.app/',
            changeOrigin: true,
        })
    );
    app.use(
        '/file',
        createProxyMiddleware({
            target: 'https://me-server-7t4azswppa-du.a.run.app/',
            changeOrigin: true,
        })
    );
};