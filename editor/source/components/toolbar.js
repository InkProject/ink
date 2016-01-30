import React from 'react'
import Component from './index'
import classNames from 'classnames'

import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { toolbarAction } from '../actions'

class Toolbar extends Component {
    constructor(props) {
        super(props)
        this.state = {confirm: false}
    }
    showConfirm(flag) {
        this.setState({confirm: flag})
    }
    onRemoveClick() {
        this.props.toolbarAction.removeArticle()
        this.setState({confirm: false})
    }
    onSaveClick() {
        this.props.toolbarAction.saveContent()
    }
    render() {
        return (
            <ul id="right">
                <li><a className="button button-cube"><i className="fa fa-rocket"></i>发布</a></li>
                <li><a className="button button-cube deploy" href="/" target="_blank"><i className="fa fa-chrome"></i>预览</a></li>
                <li><button className="button button-circle" onClick={() => this.onSaveClick()}><i className="fa fa-floppy-o"></i></button></li>
                <li><button className="button button-circle remove" onFocus={() => this.showConfirm(true)} onBlur={() => this.showConfirm(false)}><i className="fa fa-trash"></i></button></li>
                {this.state.confirm ? <div id="confirm" className="hover" onMouseDown={() => this.onRemoveClick()}>确认删除</div> : null}
            </ul>
        )
    }
}

export default connect(null, function(dispatch) {
    return {
        toolbarAction: bindActionCreators(toolbarAction, dispatch)
    }
})(Toolbar)
