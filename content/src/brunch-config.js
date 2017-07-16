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
    npm: {
        enabled: true,
        aliases: {
            "gocms/base/components/gForm/GForm": 'gocms-brunch-slipstream',
            "gocms/base/components/gForm/GInput": 'gocms-brunch-slipstream',
            "gocms/base/components/gForm/GTextArea": 'gocms-brunch-slipstream',
            "gocms/base/services/api": 'gocms-brunch-slipstream',
            "gocms/base/actions/apiRequestActions": 'gocms-brunch-slipstream'
        }
    },
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
        copycat: {},
        gocmsSlipstream: {
            enabled: true
        }
    },
    modules: {
        nameCleaner: path => path.replace('', "contact-form.plugins.gocms.io/")
    },
    optimize: false,
    notifications: false,
    hot: false,
    paths: {
        public: '../',
        watched: [
            // 'admin',
            'public',
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
