import React from 'react'
import Component from '../index'
import classNames from 'classnames'
import moment from 'moment'
moment.locale('zh-cn')
import _ from 'lodash'

import { connect } from 'react-redux'
import { Link } from 'react-router'

import * as listAction from './action'

class List extends Component {
    constructor(props) {
        super(props)
        this.state = {
            selected: 'article',
            filterObj: {}
        }
    }
    componentWillMount() {
        store.dispatch(listAction.fetch())
    }
    componentDidUpdate() {
        this.refs.search.focus()
    }
    onClickList() {
        store.dispatch(listAction.hide())
        this.clearSeacrh()
    }
    onClickNav(item, event) {
        this.setState({
            selected: item
        })
        event.stopPropagation()
    }
    onSearch() {
        const keyword = _.trim(this.refs.search.value).toLowerCase()
        const searchObj = this.props.list.get('searchObj')
        let filterObj = {}
        _.each(searchObj, (item, id) => {
            if (item.indexOf(keyword) < 0) {
                filterObj[id] = true
            }
        })
        this.setState({
            filterObj
        })
    }
    setSeacrh(tag, event) {
        this.refs.search.value = tag
        this.onSearch()
        this.setState({
            selected: 'article'
        })
        event.stopPropagation()
    }
    clearSeacrh() {
        this.refs.search.value = ''
        this.onSearch()
    }
    render() {
        const { list, editor } = this.props
        const countObj = _.countBy(list.get('data').toJS(), 'draft')
        const draftCount = countObj['true'] || 0
        const nonDraftCount = countObj['false'] || 0
        const tagCount = Object.keys(list.get('tags')).length
        const hideEmpty = nonDraftCount && this.state.selected == 'article' ||
          draftCount && this.state.selected == 'draft' || tagCount && this.state.selected == 'tags'
        return (
            <div id="list-wrap" onClick={this.onClickList.bind(this)} className={classNames({hide: !list.get('show')})}>
                <img src={require('../../../assets/logo.png')} className="logo" />
                <label className="search-wrap" htmlFor="search">
                    <i className="fa fa-search"></i>
                <input ref="search" id="search" type="text" placeholder="关键字搜索..." onChange={this.onSearch.bind(this)} />
                </label>
                <div className={classNames('list', {[this.state.selected]: true})}>
                    <ul className="list-nav">
                        <li className={classNames('item', 'article', {selected: this.state.selected == 'article'})} onClick={this.onClickNav.bind(this, 'article')}><i className="fa fa-paper-plane-o"></i>文章<span className="count">({nonDraftCount})</span></li>
                        <li className={classNames('item', 'draft', {selected: this.state.selected == 'draft'})} onClick={this.onClickNav.bind(this, 'draft')}><i className="fa fa-inbox"></i>草稿<span className="count">({draftCount})</span></li>
                    <li className={classNames('item', 'tags', {selected: this.state.selected == 'tags'})} onClick={this.onClickNav.bind(this, 'tags')}><i className="fa fa-tags"></i>标签<span className="count">({tagCount})</span></li>
                    </ul>
                    <i className={classNames('empty', 'fa', 'fa-inbox', {hide: hideEmpty})}></i>
                    <ul className={classNames('list-content', {hide: this.state.selected == 'tags'})}>{
                        list.get('data').map(item => {
                            const date = moment(item.get('date')).fromNow()
                            return <Link to={`/edit/${item.get('id')}`} className={classNames({hide: this.state.filterObj[item.get('id')]})}>
                                <li className={classNames('item hover', {selected: item.get('id') == editor.get('id'), draft: item.get('draft'), article: !item.get('draft')})} key={item.get('id')}>
                                <div className="title">{item.get('title')}</div>
                                <div className="date">{date}</div>
                                </li>
                            </Link>
                        })
                    }</ul>
                    <ul className={classNames('tags-content', {hide: this.state.selected != 'tags'})}>{
                        Object.keys(list.get('tags')).map((tag, count) => {
                            return <li className="item" onClick={this.setSeacrh.bind(this, tag)}>{tag}</li>
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
