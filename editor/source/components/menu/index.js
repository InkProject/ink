import React from 'react'
import Component from '../index'
import classNames from 'classnames'
import { connect } from 'react-redux'
import { Link } from 'react-router'

import * as listAction from '../list/action'
import * as menuAction from '../menu/action'
import * as util from '../util'

import _ from 'lodash'

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
    onFullscreen() {
        if ((document.fullScreenElement && document.fullScreenElement !== null) ||
            (!document.mozFullScreen && !document.webkitIsFullScreen)) {
            if (document.documentElement.requestFullScreen) {
                document.documentElement.requestFullScreen()
            } else if (document.documentElement.mozRequestFullScreen) {
                document.documentElement.mozRequestFullScreen()
            } else if (document.documentElement.webkitRequestFullScreen) {
                document.documentElement.webkitRequestFullScreen(Element.ALLOW_KEYBOARD_INPUT)
            }
            _.delay(() => {
                store.dispatch(util.showTip('auto', '缩放页面以获得最佳视觉体验'))
            }, 5000)
        } else {
            if (document.cancelFullScreen) {
                document.cancelFullScreen()
            } else if (document.mozCancelFullScreen) {
                document.mozCancelFullScreen()
            } else if (document.webkitCancelFullScreen) {
                document.webkitCancelFullScreen()
            }
        }
    }
    render() {
        const { show } = this.props.list.toJS()
        const { focusMode } = this.props.menu.toJS()
        const path = this.props.routing.location.pathname
        return (
            <div id="menu" className={classNames({close: !show})}>
                <ul className="tool">
                    <li className="inbox">
                        <button className="button button-circle" onClick={this.onInboxClick.bind(this)}>
                            <i className={classNames('fa', {['fa-'+(show?'plus':'inbox')]: true})}></i>
                        </button>
                    </li>
                    <li className="focus" onClick={this.onChangeFocusMode.bind(this)}>
                        <button className={classNames('button', 'button-circle')}>
                            <i className={classNames('fa', {'fa-dot-circle-o': focusMode, 'fa-circle-o': !focusMode})}></i>
                        </button>
                    </li>
                    <li className="fullscreen" onClick={this.onFullscreen.bind(this)}><button className="button button-circle"><i className="fa fa-crop"></i></button></li>
                    <li className="theme"><button className="button button-circle"><i className="fa fa-moon-o"></i></button></li>
                    <li className="setting">
                        <Link to="/edit/config">
                            <button className={classNames('button', 'button-circle', {'active': path == '/edit/config'})}><i className="fa fa-wrench"></i></button>
                        </Link>
                    </li>
                    <li className="help">
                        <Link to="/edit/help">
                            <button className={classNames('button', 'button-circle', {'active': path == '/edit/help'})}><i className="fa fa-hashtag"></i></button>
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
