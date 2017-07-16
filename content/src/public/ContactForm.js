import React from 'react'
import {contact_form_submit} from './config/actions';
import {connect} from 'react-redux'
import GForm from "gocms/base/components/gForm/GForm"
import GInput from 'gocms/base/components/gForm/GInput'
import GTextArea from 'gocms/base/components/gForm/GTextArea'

class ContactForm extends React.Component {
    constructor(props) {
        super(props);
        this.handleSubmit = this.handleSubmit.bind(this); //bind function once
        this.state = {
            fields: {
                name: {
                    label: this.props.nameLabel || "Name",
                },
                email: {
                    label: this.props.emailLabel || "Email",
                },
            },
            submit: {
                btnLabel: this.props.submitBtnLabel || "Submit",
                btnClassName: this.props.submitBtnClassName || "",
                onSubmit: this.handleSubmit,
            },
            shake: false,
            errMessage: this.props.errMessage || null
        };
    }

    componentWillReceiveProps(nextProps) {
        if (!!nextProps.err && nextProps.reqTime != this.props.reqTime) {
            this.setState({shake: true})
        }
        if (!!nextProps.errMessage && nextProps.errMessage != this.state.errMessage) {
            this.setState({errMessage: nextProps.errMessage});
        }
    }

    handleSubmit(model) {
        this.props.contact_form_submit(this.props.name, model);
        this.setState({errMessage: null});
    }

    componentDidMount() {

    }

    render() {

        return (
            <GForm id={this.props.id} className={this.props.className} name={this.props.name}
                   onSubmit={this.state.submit.onSubmit}
                   submitBtn={this.state.submit.btnLabel}
                   submitBtnClassName={this.state.submit.btnClassName}
                   submitBtnShake={this.state.shake}>
                <GInput id="name" name="name" type="text" label={this.state.fields.name.label} required/>
                <GInput id="email" name="email" type="text" label={this.state.fields.email.label} validations="isEmail"
                        validationError="Please enter a valid email." required/>
                {this.props.children}
            </GForm>
        )
    }
}

function mapStateToProps(state, ownProps) {
    let errMessage;
    let err;
    let reqTime;
    let req = state.api.request[ownProps.name];
    if (!!req) {
        reqTime = req.receivedAt;
        if (!!req.err) {
            err = req.err;
            if (!!err.json && !!err.json.message) {
                errMessage = err.json.message;
            }
        }
    }
    return {
        reqTime: reqTime,
        err: err,
        errMessage: errMessage,
    }
}

export default connect(mapStateToProps, {
    contact_form_submit
})(ContactForm);