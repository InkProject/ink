import React from 'react'
import Component from '../index'
import classNames from 'classnames'

import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import * as menuAction from '../menu/action'

class Welcome extends Component {
    render() {
        return (
            <div className="welcome-wrap">
                <img src={require('../../../assets/logo.png')} className="logo" />
                <div className="slogan">构建只为纯粹书写的博客</div>
                <div className="guide">
                    <div className="document hover"><i className="fa fa-hashtag"></i>撰写指南</div>
                    <div className="getstart hover" onClick={() => this.props.menuAction.showModal(true)}><i className="fa fa-circle-o"></i>创建文章</div>
                </div>
            </div>
        )
    }
}

export default connect(null, function(dispatch) {
    return {
        menuAction: bindActionCreators(menuAction, dispatch)
    }
})(Welcome)
