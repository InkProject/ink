import 'font-awesome/css/font-awesome.css'
import './styles/index.css'

import React from 'react'
import Component from './components'
import ReactDom from 'react-dom'
import classNames from 'classnames'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { editorAction } from './actions'

import Left from './components/left'
import Editor from './components/editor'
import Search from './components/search'
import Toolbar from './components/toolbar'
import Modal from './components/modal'
import Tooltip from './components/tooltip'

class App extends Component {
    render() {
        const list = this.props.list
        const util = this.props.util
        const editor = this.props.editor
        const tip = util.get('tip')
        return (
            <div id="container">
                <Modal />
                <Tooltip />
                <Left />
                <Search />
                <Toolbar />
                { this.props.Welcome || this.props.Editor }
            </div>
        )
    }
}

export default connect(function(state) {
    return state
})(App)
