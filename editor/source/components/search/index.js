import React from 'react'
import Component from '../index'
import classNames from 'classnames'

import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import * as listAction from '../list/action'

export default class Search extends Component {
    constructor(props) {
        super(props)
    }
    onFocus() {
        store.dispatch(listAction.show())
    }
    render() {
        const { show } = this.props.list.toJS()
        return (
            <label id="search-wrap" htmlFor="search" className={classNames({focus: show})}>
                <i className="fa fa-search"></i>
                <input id="search" type="text" placeholder="搜索..." onFocus={() => this.onFocus()} onBlur={() => this.onBlur()} />
            </label>
        )
    }
}

export default connect(function(state) {
    return {
        list: state.list
    }
})(Search)
