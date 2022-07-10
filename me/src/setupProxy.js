const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
    app.use(
        '/v2',
        createProxyMiddleware({
            target: process.env.SERVER_IP,
            changeOrigin: true,
        })
    );
};