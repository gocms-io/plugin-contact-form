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

/////////////////////////////////////////////////////////////////
// list all components that need to be included in compile //////
/////////////////////////////////////////////////////////////////
// require("../ContactFormAdmin"); // this is not needed because it is hotloaded via a route