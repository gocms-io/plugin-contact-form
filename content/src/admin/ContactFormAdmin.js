import React from 'react'
import {connect} from 'react-redux'

class ContactFormAdmin extends React.Component {
    constructor(props) {
        super(props);
    }

    componentWillReceiveProps(nextProps) {

    }


    componentDidMount() {

    }

    render() {
        return (
            <div className="contact-form-plugins-gocms-admin-placeholder">
                <h1>Contact Form Admin</h1>
                <p>Placeholder for the contact form plugin admin section.</p>
            </div>
        )
    }
}

function mapStateToProps(state, ownProps) {

    return {}
}

export default connect(mapStateToProps, {})(ContactFormAdmin);