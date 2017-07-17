export const CONTACT_FORM_SUBMIT = "CONTACT_FORM_SUBMIT";

export function contact_form_submit(key, data) {
    console.log("Data: ", data);
    return {
        type: CONTACT_FORM_SUBMIT,
        key,
        data,
        requestedAt: Date.now()
    }
}