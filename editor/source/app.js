import 'font-awesome/less/font-awesome.less';
import './styles/index.less';

import React from 'react';
import ReactDom from 'react-dom';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import classNames from 'classnames';
import actions from './actions';

import Editor from './components/editor';
import List from './components/list';
import Menu from './components/menu';
import Header from './components/header';
import Search from './components/search';
import Toolbar from './components/toolbar';

class App extends React.Component {
    componentWillMount() {
        this.props.actions.fetchList();
    }
    render() {
        const list = this.props.list;
        const actions = this.props.actions;
        const content = this.props.content;
        return (
            <div id="container">
                <div id="left" onMouseLeave={actions.hideList} onMouseOver={actions.showList}>
                    <Menu {...this.props} />
                    <div id="files" className={classNames({hide: !list.get('show')})}>
                        <Search {...this.props} />
                        <List {...this.props} />
                    </div>
                </div>
                <div id="right">
                    <Toolbar {...this.props} />
                </div>
                <Header {...this.props} />
                <Editor {...this.props} />
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
