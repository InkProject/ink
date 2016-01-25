import React from 'react';
import Component from './index';
import classNames from 'classnames';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { toolbarAction } from '../actions';

class Welcome extends Component {
    render() {
        return (
            <div className="welcome-wrap">
                <img src={require('../../assets/logo.png')} className="logo" />
                <div className="slogan">构建只为纯粹书写的博客</div>
            </div>
        );
    }
}

export default connect(null, function(dispatch) {
    return {
        // toolbarAction: bindActionCreators(toolbarAction, dispatch)
    };
})(Welcome);