// @ts-nocheck

import { Button, Input, Text, Link, Grid, Spacer, Card, FormElement } from '@nextui-org/react';
import { Flex } from '../styles/flex';
import { HouseIcon } from '../icons/breadcrumb/house-icon';
import { UsersIcon } from '../icons/breadcrumb/users-icon';
import { ChangeEvent, useState } from 'react';
import { createAppAPI } from '../../api/create-app';
import router from 'next/router';

export const AddModel = () => {
    const [values, setValues] = useState({
        appName: '',
        dockerImage: '',
    })

    const handleChange = (prop: string) => (event: ChangeEvent<FormElement>) => {
        setValues({ ...values, [prop]: event.target.value })
    }
    const createApp = async () => {
        try {
            const response = await createAppAPI(values.appName, values.dockerImage);
            console.log(response);
            if (response.message === 'success') {
                router.push('/list')
            }
        } catch (error) {
            console.log(error);
        }
    };
    return (
        <Flex
            css={{
                'mt': '$5',
                'px': '$6',
                '@sm': {
                    mt: '$10',
                    px: '$16',
                },
            }}
            justify={'center'}
            direction={'column'}
        >
            <Grid.Container gap={0} justify="flex-start">
                <Grid xs={1}>
                    <HouseIcon />
                    <Link href="/dashboard" color="text">Home</Link>
                    <Text>&nbsp;/</Text>
                </Grid>
                <Grid xs={1}>
                    <UsersIcon />
                    <Link href="#" color="text">Model</Link>
                    <Text>&nbsp;/</Text>
                </Grid>
                <Grid xs={1}>
                    <Link href="#" color="text">Submit</Link>
                </Grid>
            </Grid.Container>
            <Spacer></Spacer>
            <Card>
                <Card.Body>
                    <Text h3>Add Model</Text>
                    <Spacer y={3}></Spacer>
                    <Grid.Container gap={2} justify="center">
                        <Grid xs={2}>
                        </Grid>
                        <Grid xs={8} css={{ justifyContent: "center" }}>
                            <Input
                                underlined
                                labelLeft="Model Name"
                                placeholder="GPT4"
                                width="75%"
                                color="secondary"
                                onChange={handleChange('appName')}
                                value={values.appName}
                            />
                        </Grid>
                        <Grid xs={2}>
                        </Grid>
                        <Grid xs={2}>
                        </Grid>
                        <Grid xs={8} css={{ justifyContent: "center" }}>
                            <Input
                                underlined
                                labelLeft="Docker Image"
                                placeholder="lyleshaw/gradio-test:latest"
                                width="75%"
                                color="secondary"
                                onChange={handleChange('dockerImage')}
                                value={values.dockerImage}
                            />
                        </Grid>
                        <Grid xs={2}>
                        </Grid>
                        <Grid xs={2}>
                        </Grid>
                        <Grid xs={8} css={{ justifyContent: "center" }}>
                            <Button shadow color="gradient" auto onPress={createApp}>
                                Submit
                            </Button>
                        </Grid>
                        <Grid xs={2}>
                        </Grid>
                    </Grid.Container>
                    <Spacer y={3}></Spacer>
                </Card.Body>
            </Card>
        </Flex>
    );
};
