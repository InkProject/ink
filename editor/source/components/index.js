import React from 'react';
import { immutableRenderDecorator } from 'react-immutable-render-mixin';

class Component extends React.Component {}

immutableRenderDecorator(Component);

export default Component;
