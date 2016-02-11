import React from 'react'
import Component from '../index'
import classNames from 'classnames'
import { connect } from 'react-redux'
import { Link } from 'react-router'

import * as listAction from '../list/action'
import * as menuAction from '../menu/action'
import * as util from '../util'

class Menu extends Component {
    onInboxClick() {
        if (this.props.list.get('show')) {
            store.dispatch(menuAction.showModal(true))
        } else {
            store.dispatch(listAction.show())
        }
    }
    onChangeFocusMode() {
        const focusMode = store.getState().menu.get('focusMode')
        store.dispatch(menuAction.changeFocusMode(!focusMode))
        if (!focusMode) store.dispatch(util.showTip('auto', '切换到专注模式'))
    }
    render() {
        const { show, loading } = this.props.list.toJS()
        const { focusMode } = this.props.menu.toJS()
        return (
            <ul id="menu" className="list">
                <li>
                    <button className="button button-circle inbox" onClick={this.onInboxClick.bind(this)}>
                        <i className={classNames('fa', {['fa-'+(show?'plus':'inbox')]: true})}></i>
                    </button>
                </li>
                <li onClick={this.onChangeFocusMode.bind(this)}>
                    <button className="button button-circle focus">
                        <i className={classNames('fa', {'fa-dot-circle-o': focusMode, 'fa-circle-o': !focusMode})}></i>
                    </button>
                </li>
                <li>
                    <Link to="/edit/config">
                        <button className="button button-circle setting"><i className="fa fa-wrench"></i></button>
                    </Link>
                </li>
                <li><button className="button button-circle theme"><i className="fa fa-moon-o"></i></button></li>
                <li>
                    <Link to="/edit/help">
                        <button className="button button-circle help"><i className="fa fa-hashtag"></i></button>
                    </Link>
                </li>
            </ul>
        )
    }
}

export default connect(function(state) {
    return {
        list: state.list,
        menu: state.menu
    }
})(Menu)
