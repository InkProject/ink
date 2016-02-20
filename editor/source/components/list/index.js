import React from 'react'
import Component from '../index'
import classNames from 'classnames'
import _ from 'lodash'

import { connect } from 'react-redux'
import { Link } from 'react-router'

import * as listAction from './action'

class List extends Component {
    constructor(props) {
        super(props)
        this.state = {
            selected: 'article'
        }
    }
    componentWillMount() {
        store.dispatch(listAction.fetch())
    }
    onClickList() {
        store.dispatch(listAction.hide())
    }
    onClickNav(item, event) {
        this.setState({
            selected: item
        })
        event.stopPropagation()
    }
    render() {
        const { list, editor } = this.props
        const countObj = _.countBy(list.get('data').toJS(), 'draft')
        const draftCount = countObj['true']
        const nonDraftCount = countObj['false']
        return (
            <div id="list-wrap" onClick={this.onClickList.bind(this)} className={classNames({hide: !list.get('show')})}>
                <img src={require('../../../assets/logo.png')} className="logo" />
                <label className="search-wrap" htmlFor="search">
                    <i className="fa fa-search"></i>
                <input id="search" type="text" placeholder="关键字搜索..." />
                </label>
                <div className={classNames('list', {[this.state.selected]: true})}>
                    <ul className="list-nav">
                        <li className={classNames('item', 'article', {selected: this.state.selected == 'article'})} onClick={this.onClickNav.bind(this, 'article')}><i className="fa fa-paper-plane-o"></i>文章<span className="count">({nonDraftCount})</span></li>
                        <li className={classNames('item', 'draft', {selected: this.state.selected == 'draft'})} onClick={this.onClickNav.bind(this, 'draft')}><i className="fa fa-inbox"></i>草稿<span className="count">({draftCount})</span></li>
                    <li className={classNames('item', 'tags', {selected: this.state.selected == 'tags'})} onClick={this.onClickNav.bind(this, 'tags')}><i className="fa fa-tags"></i>标签<span className="count">({Object.keys(list.get('tags')).length})</span></li>
                    </ul>
                    <ul className={classNames('list-content', {hide: this.state.selected == 'tags'})}>{
                        list.get('data').map(item => {
                            return <Link to={`/edit/${item.get('id')}`}>
                                <li className={classNames('item hover', {selected: item.get('id') == editor.get('id'), draft: item.get('draft'), article: !item.get('draft')})} key={item.get('id')}>
                                <div className="title">{item.get('title')}</div>
                                <div className="date">2015/12/11</div>
                                </li>
                            </Link>
                        })
                    }</ul>
                    <ul className={classNames('tags-content', {hide: this.state.selected != 'tags'})}>{
                        Object.keys(list.get('tags')).map((tag, count) => {
                            return <li className="item">{tag}</li>
                        })
                    }</ul>
                </div>
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
