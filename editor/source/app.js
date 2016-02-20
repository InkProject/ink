import 'font-awesome/css/font-awesome.css'
import './styles/index.css'

import React from 'react'
import Component from './components'
import ReactDom from 'react-dom'
import classNames from 'classnames'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import Menu from './components/menu'
import List from './components/list'
import Editor from './components/editor'
import Toolbar from './components/toolbar'
import Modal from './components/modal'
import Tooltip from './components/tooltip'

import * as editorAction from './components/toolbar/action'

class App extends Component {
    render() {
        const list = this.props.list
        const util = this.props.util
        const editor = this.props.editor
        const tip = util.get('tip')
        // this.props.editorAction.reset()
        return (
            <div id="container">
                <Modal />
                <Tooltip />
                <Menu />
                <List />
                <Toolbar />
                { this.props.Welcome || this.props.Editor }
            </div>
        )
    }
}

export default connect(function(state) {
    return state
})(App)
