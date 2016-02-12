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
        const { show } = this.props.list.toJS()
        const { focusMode } = this.props.menu.toJS()
        const path = this.props.routing.location.pathname
        return (
            <div id="menu" className={classNames({close: !show})}>
                <ul className="head">
                    <li onClick={this.onChangeFocusMode.bind(this)}>
                        <button className={classNames('button', 'button-circle', 'focus')}>
                            <i className={classNames('fa', {'fa-dot-circle-o': focusMode, 'fa-circle-o': !focusMode})}></i>
                        </button>
                    </li>
                    <li><button className="button button-circle theme"><i className="fa fa-moon-o"></i></button></li>
                </ul>
                <ul className="tool">
                    <li>
                        <button className="button button-circle inbox" onClick={this.onInboxClick.bind(this)}>
                            <i className={classNames('fa', {['fa-'+(show?'plus':'inbox')]: true})}></i>
                        </button>
                    </li>
                    <li>
                        <Link to="/edit/config">
                            <button className={classNames('button', 'button-circle', 'setting', {'active': path == '/edit/config'})}><i className="fa fa-wrench"></i></button>
                        </Link>
                    </li>
                    <li>
                        <Link to="/edit/help">
                            <button className={classNames('button', 'button-circle', 'help', {'active': path == '/edit/help'})}><i className="fa fa-hashtag"></i></button>
                        </Link>
                    </li>
                </ul>
            </div>
        )
    }
}

export default connect(function(state) {
    return {
        list: state.list,
        menu: state.menu,
        routing: state.routing
    }
})(Menu)
