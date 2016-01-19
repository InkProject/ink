import 'font-awesome/css/font-awesome.css';
import './styles/index.css';

import React from 'react';
import Component from './components';
import ReactDom from 'react-dom';
import classNames from 'classnames';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { editorAction } from './actions';

import Left from './components/left';
import Editor from './components/editor';
import Search from './components/search';
import Toolbar from './components/toolbar';

class App extends Component {
    render() {
        const list = this.props.list;
        const util = this.props.util;
        const tip = util.get('tip');
        return (
            <div id="container">
                <div id="tooltip" className={classNames({hide: !tip.get('show'), error: tip.get('error')})}>
                    <div className="content">
                        {tip.get('loading') ? <i className="fa fa-cog fa-spin"></i> : null}
                        {tip.get('content')}
                    </div>
                </div>
                <Left list={list} listComponent={this.props.List} />
                <Search />
                <Toolbar />
                {this.props.Editor}
            </div>
        );
    }
}

export default connect(function(state) {
    return state;
})(App);
