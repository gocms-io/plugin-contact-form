module.exports = {
    sourceMaps: 'inline',
    files: {
        javascripts: {
            joinTo: {
                "public.js": [/^public/, "initialize.js"],
                // "admin.js": /^admin/,
                "public_vendor.js": /^(?!public|admin)/,
            }
        },
        stylesheets: {
            joinTo: {
                'public.css': ["public/styles/index.scss"],
                // 'admin.css': ["admin/config/styles/index.scss"],
                // 'admin_ie.css': ["app/config/styles/ie.scss"],
            }
        }
    },
    npm: {},
    plugins: {
        babel: {
            presets: ['es2015', 'stage-0', 'react'],
            ignore: [/^(node_modules|bower_components|vendor)/],
            plugins: [
                ['transform-runtime', 'transform-object-assign', {
                    polyfill: false,
                    regenerator: true
                },
                ]
            ]
        },
        sass: {
            options: {
                includePaths: [],
                precision: 15
            },
            mode: 'native',
            sourceMapEmbed: true,
            debug: 'comments'
        },
        copycat: {}
    },
    modules: {
        nameCleaner: path => path.replace(/^(public|admin)/, "contact-form.plugins.gocms.io")
    },
    optimize: false,
    notifications: false,
    hot: false,
    paths: {
        public: '../',
        watched: [
            // 'admin',
            'public',
            'public_vendor',
            'initialize.js'
        ]
    },
    overrides: {
        production: {
            optimize: true,
            sourceMaps: false,
            plugins: {
                autoReload: {
                    enabled: false
                }
            }
        }
    }
};
