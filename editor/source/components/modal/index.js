import React from 'react'
import Component from '../index'
import classNames from 'classnames'
import _ from 'lodash'

import { connect } from 'react-redux'
import * as modalAction from './action'
import * as menuAction from '../menu/action'

class Modal extends Component {
    constructor(props) {
        super(props)
    }
    closeModal() {
        store.dispatch(menuAction.showModal(false))
    }
    createArticle() {
        const name = _.trim(this.input.value)
        if (name) {
            this.input.value = ''
            store.dispatch(modalAction.createArticle(name))
            store.dispatch(menuAction.showModal(false))
        } else {
            this.input.focus()
        }
    }
    render() {
        const { menu } = this.props
        return (
            <div id="modal" className={classNames({hide: !menu.get('modal').get('show')})}>
                <div className="overlay" onClick={() => this.closeModal()}></div>
                <div className="content">
                    <div className="title"><i className="fa fa-inbox"></i>创建新文章</div>
                    <input type="text" className="name" placeholder="文件名，建议使用-与字母" ref={(input) => {
                        if (input) {
                            input.focus()
                            input.setSelectionRange(0, input.value.length)
                            this.input = input
                        }
                    }} />
                <div className="confirm hover" onClick={() => this.createArticle()}><i className="fa fa-inbox"></i>确认</div>
                </div>
            </div>
        )
    }
}

export default connect(function(state) {
    return {
        menu: state.menu
    }
})(Modal)
