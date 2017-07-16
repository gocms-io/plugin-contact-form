import {fork,takeEvery, put, call} from 'redux-saga/effects';
import * as contactFormActions from './actions';
import * as apiActions from 'gocms/base/actions/apiRequestActions';
import {Post, ENDPOINTS} from 'gocms/base/services/api';

// watch for login requests
function* watchContactFormSaga() {
    yield takeEvery(contactFormActions.CONTACT_FORM_SUBMIT, handleContactFormSaga); // see details what is REQUEST param below
}

function* handleContactFormSaga(action) {
    let {res, err} = yield call(Post, "/api/contact-form.plugins.gocms.io/contact-form", action.data); // calling our api method
    if (res) {
        // push user info to store and storage
        yield put(apiActions.success(action.key, res.data));
    }
    else if (err) {
        yield put(apiActions.failure(action.key, err));
    }
}

/**
 * rootSaga
 */
export default function* rootSaga() {
    yield [
        fork(watchContactFormSaga),
    ];
}
