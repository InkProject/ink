import React from 'react'
import Component from '../index'
import classNames from 'classnames'

import { connect } from 'react-redux'
import { Link } from 'react-router'

import * as listAction from './action'

class List extends Component {
    componentWillMount() {
        store.dispatch(listAction.fetch())
    }
    render() {
        const { list, editor } = this.props
        return (
            <ul id="list" className={classNames({hide: !list.get('show')})}>{
                list.get('data').map(item => {
                    return <li className={classNames('item hover', {selected: item.get('id') == editor.get('id')})} key={item.get('id')}>
                        <Link to={`/edit/${item.get('id')}`}>
                            <div className="head">
                                <span className="date">2015-03-01 18:00</span>
                            </div>
                            <div className="name">{item.get('name')}</div>
                            <div className="title">{item.get('title')}</div>
                            <div className="preview">{item.get('preview')}</div>
                        </Link>
                    </li>
                })
            }</ul>
        )
    }
}

export default connect(function(state) {
    return {
        list: state.list,
        editor: state.editor
    }
})(List)
