import React from 'react'
import {contact_form_submit} from './config/actions';
import {connect} from 'react-redux'
import CSSTransitionGroup from 'react-transition-group/CSSTransitionGroup'
import GForm from "gocms/base/components/gForm/GForm"
import GInput from 'gocms/base/components/gForm/GInput'
import GError from 'gocms/base/components/gForm/GError'

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
                onSubmit: this.handleSubmit
            },
            busy: false,
            shake: false,
            errMessage: this.props.errMessage || null,
            success: false
        };
    }

    componentWillReceiveProps(nextProps) {
        if (nextProps.reqTime != this.props.reqTime) {
            // busy button
            this.setState({busy: false});

            // err shake button
            if (!!nextProps.err) {
                this.setState({shake: true});
            }
            // good request display success
            else {
                this.setState({success: true})
            }
        }

        // error message
        if (!!nextProps.errMessage && nextProps.errMessage != this.state.errMessage) {
            this.setState({errMessage: nextProps.errMessage});
        }


    }

    handleSubmit(model) {
        this.props.contact_form_submit(this.props.name, model);
        this.setState({
            errMessage: null,
            busy: true
        });
    }

    componentDidMount() {

    }

    render() {
        let html = null;
        //todo figureout why animation doesn't work
        // todo then finish the admin injection for plugin frontend
        // todo then finish remaining contact forms
        // todo then finish about page
        // if the button shakes stop it!
        setTimeout(function () {
            this.setState({success: true});
        }.bind(this), 1000);

        if (this.state.success) {
            html =
                <div dangerouslySetInnerHTML={{__html: this.props.successMsg}}/>
            ;
        }
        else {
            html =
                <GForm id={this.props.id} className={this.props.className} name={this.props.name}
                       onSubmit={this.state.submit.onSubmit}
                       submitBtn={this.state.submit.btnLabel}
                       submitBtnClassName={this.state.submit.btnClassName}
                       submitBtnShake={this.state.shake}
                       submitBtnBusy={this.state.busy}
                >
                    <GInput id="name" name="name" type="text" label={this.state.fields.name.label} required/>
                    <GInput id="email" name="email" type="text" label={this.state.fields.email.label}
                            validations="isEmail"
                            validationError="Please enter a valid email." required/>
                    {this.props.children}
                    <GError
                        errMessage={this.state.errMessage}
                    />
                </GForm>
            ;
        }
        return (
            <CSSTransitionGroup transitionName="contact-form-plugins-gocms-io-animation"
                                transitionEnterTimeout={500}
                                transitionLeaveTimeout={500}>
                {html}
            </CSSTransitionGroup>
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