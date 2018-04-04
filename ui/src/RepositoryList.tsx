import * as React from 'react';

import { Button, List } from 'semantic-ui-react';

// import {Renderer} from 'react-dom'

export class RepositoryList extends React.Component {

    state: { list: string[] };
    constructor(props: any) {
        super(props);
        this.state = { list: [] };
        this.props = props;
    }

    componentDidMount() {
        console.log('fuck!!!');
        fetch('http://localhost:8080/api/repositories').then(
            res => {
                res.json().then(j => {
                    console.log(j);
                    this.setState({ list: j });

                });
            }
        );
    }

    render() {
        console.log(this, 'is rendered');

        return (
            <List divided verticalAlign="middle">
                {
                    this.state.list.map(n => (
                        <List.Item key={n}>
                            <List.Content floated="right">
                                <Button>Info</Button>
                            </List.Content>
                            {/* <Image avatar src='/assets/images/avatar/small/lena.png' /> */}
                            <List.Content>
                                {n}
                            </List.Content>
                        </List.Item>
                    )
                    )
                }
            </List>
        );
    }
}