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
            <div id="list-wrap" onClick={() => { store.dispatch(listAction.hide()) }} className={classNames({hide: !list.get('show')})}>
                <img src={require('../../../assets/logo.png')} className="logo" />
                <ul className="list-nav">
                    <li className="item selected"><i className="fa fa-paper-plane-o"></i>已发布</li>
                    <li className="item"><i className="fa fa-inbox"></i>草稿</li>
                    <li className="item tags"><i className="fa fa-tags"></i>标签</li>
                </ul>
                <ul id="list">{
                    list.get('data').map(item => {
                        return <li className={classNames('item hover', {selected: item.get('id') == editor.get('id')})} key={item.get('id')}>
                            <Link to={`/edit/${item.get('id')}`}>
                                <div className="title">{item.get('title')}</div>
                                <div className="preview">{item.get('preview')}</div>
                                <div className="footer">
                                    <div className="name">{item.get('name')}</div>
                                    <div className="date">2015-03-01 18:00</div>
                                    <div className="tags">{
                                        item.get('tags').map(item =>
                                            item ? <span className="tag">{item}</span> : null
                                        )
                                    }</div>
                                </div>
                            </Link>
                        </li>
                    })
                }</ul>
            </div>
        )
    }
}

export default connect(function(state) {
    return {
        list: state.list,
        editor: state.editor
    }
})(List)
