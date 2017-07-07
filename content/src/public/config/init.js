'use strict';

import sagas from './sagas';
import routes from './routes';
import reducers from './reducers'

const plugin = {
    id: "contact-form.plugins.gocms.io",
    name: "contact form",
    sagas: sagas,
    routes: routes,
    reducers: reducers,
};

export default plugin;