import React from 'react'
import Component from '../index'
import classNames from 'classnames'

import { connect } from 'react-redux'

import * as listAction from '../list/action'
import * as menuAction from '../menu/action'

class Menu extends Component {
    onInboxClick() {
        if (this.props.show) {
            store.dispatch(menuAction.showModal(true))
        } else {
            store.dispatch(listAction.show())
        }
    }
    render() {
        const { show, loading } = this.props
        return (
            <ul id="menu" className="list">
                <li>
                    <button className="button button-circle inbox" onClick={() => this.onInboxClick()}>
                        <i className={classNames('fa', {['fa-'+(show?'plus':'inbox')]: true})}></i>
                    </button>
                </li>
                <li><button className="button button-circle focus"><i className="fa fa-unlock-alt"></i></button></li>
                <li><button className="button button-circle setting"><i className="fa fa-wrench"></i></button></li>
                <li><button className="button button-circle theme"><i className="fa fa-moon-o"></i></button></li>
                <li><button className="button button-circle help"><i className="fa fa-hashtag"></i></button></li>
                <li><i id="loading" className={classNames('fa fa-cog fa-spin', {hide: !loading})}></i></li>
            </ul>
        )
    }
}

export default connect(null)(Menu)
