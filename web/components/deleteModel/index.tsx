// @ts-nocheck

import { useState } from 'react';
import {
    Modal,
    Button,
    Text,
    Spacer,
    Grid
} from '@nextui-org/react';
import { toast } from 'react-toastify';
import { IconButton } from '../table/table.styled';
import { deleteAppAPI } from '../../api/delete-app';
import { DeleteIcon } from '../icons/table/delete-icon';

interface Props {
    setReloading: any,
    modelID: any;
}

export const DeleteModel = ({ setReloading, modelID }: Props) => {
    console.log('props', modelID)
    const [visible, setVisible] = useState(false);
    const handler = () => setVisible(true);
    const closeHandler = () => {
        setVisible(false);
    }
    const deleteApp = async () => {
        try {
            const response = await deleteAppAPI(modelID);
            console.log(response);
            if (response.message === 'success') {
                setReloading(true);
                setVisible(false);
            }
        } catch (error) {
            console.log(error);
            toast.error('失败，请重试！');
        }
    };

    return (
        <div>
            <IconButton
                onClick={() => handler()}
            >
                <DeleteIcon size={20} fill="#FF0080" />
            </IconButton>
            <Modal
                closeButton
                blur
                aria-labelledby="modal-title"
                open={visible}
                onClose={closeHandler}
            >
                <Modal.Header>
                    <Text
                        h2
                        color='error'
                        transform='uppercase'
                        id="modal-title"
                    >
                        Are You Sure to Delete This Model?
                    </Text>
                </Modal.Header>
                <Modal.Body>
                    <Grid.Container gap={4} justify="flex-start">
                        <Grid xs={6}>
                            <Button auto flat size='lg' bordered color='gradient' onPress={closeHandler}>
                                Nope
                            </Button>
                        </Grid>
                        <Grid xs={6} css={{ justifyContent: "center" }}>
                            <Button auto shadow size='lg' color='error' onPress={deleteApp}>
                                Yeah
                            </Button>
                        </Grid>
                    </Grid.Container>
                    <Spacer y={3}></Spacer>
                </Modal.Body>
                <Modal.Footer>
                </Modal.Footer>
            </Modal>
        </div>
    );
};
