import 'font-awesome/less/font-awesome.less';
import './styles/index.less';

import React from 'react';
import ReactDom from 'react-dom';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import classNames from 'classnames';

import Editor from './components/editor';

import * as actions from './actions';

class App extends React.Component {
    componentWillMount() {
        this.props.actions.fetchList();
    }
    render() {
        const actions = this.props.actions;
        const list = this.props.list;
        const content = this.props.content;
        return (
            <div id="container">
                <div id="left" onMouseLeave={actions.hideList}>
                    <ul className="menu">
                        <li className="button button-circle create" onMouseOver={actions.showList}>
                            <i className={classNames('fa', {['fa-'+(list.get('show')?'plus':'inbox')]: true})}></i>
                        </li>
                        <li className="button button-circle setting"><i className="fa fa-wrench"></i></li>
                        <li className="button button-circle"><i className="fa fa-hashtag"></i></li>
                        <li><i id="loading" className={classNames('fa fa-cog fa-spin', {hide: !list.get('loading')})}></i></li>
                    </ul>
                    <div id="files" className={classNames({hide: !list.get('show')})}>
                        <label id="search-wrap" htmlFor="search">
                            <i className="fa fa-search"></i>
                            <input id="search" type="text" placeholder="搜索..." />
                        </label>
                        <ul id="list">{
                            list.get('data').map(item =>
                                <li className="item hover" key={item.get('id')} onClick={actions.openArticle.bind(this, item.get('id'))}>
                                    <div className="head">
                                        <span className="date">2015-03-01 18:00</span>
                                    </div>
                                    <div className="title">{item.get('title')}</div>
                                    <div className="preview">{item.get('preview')}</div>
                                </li>
                            )
                        }</ul>
                    </div>
                </div>
                <div id="right">
                    <ul className="menu">
                        <li className="button button-cube"><i className="fa fa-rocket"></i>发布</li>
                        <li className="button button-cube deploy"><i className="fa fa-chrome"></i>预览</li>
                        <li className="button button-circle"><i className="fa fa-floppy-o"></i></li>
                        <li className="button button-circle remove"><i className="fa fa-trash"></i></li>
                    </ul>
                </div>
                <div id="header">
                    <div className="title">构建只为纯粹书写的博客</div>
                    <div className="info">
                        <span className="draft">草稿</span>
                        <span className="tag">产品</span>
                        <span className="tag">设计</span>
                    </div>
                </div>
                <Editor content={content}></Editor>
            </div>
        );
    }
}

export default connect(function(state) {
    return state;
}, function(dispatch) {
    return {
        actions: bindActionCreators(actions, dispatch)
    };
})(App);
