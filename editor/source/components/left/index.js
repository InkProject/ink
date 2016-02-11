import React from 'react'
import Component from '../index'
import classNames from 'classnames'
import List from '../list'
import Menu from '../menu'
import { connect } from 'react-redux'

export default class Left extends Component {
    render() {
        const { list } = this.props
        return (
            <div id="left" className={classNames({close: !list.get('show')})}>
                <Menu />
                <List />
            </div>
        )
    }
}

export default connect(function(state) {
    return {
        list: state.list
    }
})(Left)
