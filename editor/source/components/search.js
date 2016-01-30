import React from 'react'
import Component from './index'
import classNames from 'classnames'

import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { listAction } from '../actions'

export default class Search extends Component {
    constructor(props) {
        super(props)
        this.state = {focus: false}
    }
    onFocus() {
        this.props.listAction.showList()
        this.setState({focus: true})
    }
    onBlur() {
        this.setState({focus: false})
    }
    render() {
        return (
            <label id="search-wrap" htmlFor="search" className={classNames({focus: this.state.focus})}>
                <i className="fa fa-search"></i>
                <input id="search" type="text" placeholder="搜索..." onFocus={() => this.onFocus()} onBlur={() => this.onBlur()} />
            </label>
        )
    }
}

export default connect(function(state) {
    return state
}, function(dispatch) {
    return {
        listAction: bindActionCreators(listAction, dispatch)
    }
})(Search)
