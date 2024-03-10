import {SendPost} from "../../request/Api";
import React, {useEffect, useState} from "react";
import { MdDelete } from "react-icons/md"
import {
    Box,
    Button, Center,
    Divider, Flex,
    Input,
    InputGroup,
    InputRightElement,
    ListItem, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay,
    OrderedList, Spacer,
    Stack, useDisclosure
} from "@chakra-ui/react";

export function GroupListTab() {
    let [groupLists, setGroupLists] = useState([])
    let [newGroupList, setNewGroupList] = useState("")
    const [selectedGroupList, setSelectedGroupList] = useState(0)
    const { isOpen, onOpen, onClose } = useDisclosure()

    useEffect(() => {
        SendPost('api/group_lists', {}).then(r => {
            setGroupLists(r.data)
            console.log(r.data)
        })
    }, []);

    const createNewGroupList = (name) => {
        SendPost("api/new/group_lists", {name}).then(value => {
            SendPost('api/group_lists', {}).then(r => {
                setGroupLists(r.data ? r.data : [])
                console.log(r.data)
            })
        }, error => {})
    }

    const clickDeleteGroupList = (id) => {
        if (id === 0) {return}

        console.log("Delete by id: ", id)
        SendPost("api/delete/group_lists", {id}).then(value => {
            SendPost('api/group_lists', {}).then(r => {
                setGroupLists(r.data)
                console.log(r.data)
            })
        }, error => {})
    }

    return (
        <Box>
            <Stack spacing={3}>
                <InputGroup size='md'>
                    <Input
                        pr='4.5rem'
                        type={'text'}
                        placeholder='Enter new group lists'
                        value={newGroupList}
                        onChange={(event) => setNewGroupList(event.target.value)}
                    />
                    <Button colorScheme='blue' onClick={() => {createNewGroupList(newGroupList); setNewGroupList("")}}>Create</Button>
                </InputGroup>
                <OrderedList>
                    {
                        groupLists?.map((v) => <ListItem itemID={v.ID} key={v.ID}>
                            <Flex height={10}>
                                {v.Name}
                                <Spacer/>
                                <Button colorScheme='red' leftIcon={<MdDelete/>} onClick={() => {
                                    setSelectedGroupList(v.ID);
                                    onOpen();
                                }}>
                                    Delete
                                </Button>
                            </Flex>
                            <Divider/>
                        </ListItem>)
                    }
                </OrderedList>
                <Modal onClose={onClose} isOpen={isOpen} isCentered>
                    <ModalOverlay />
                    <ModalContent>
                        <ModalHeader>Are you sure?</ModalHeader>
                        <ModalCloseButton />
                        <ModalBody>
                            All admins and connect groups will be deleted!
                        </ModalBody>
                        <ModalFooter>
                            <Button onClick={onClose} mr={3}>Close</Button>
                            <Button onClick={() => {
                                clickDeleteGroupList(selectedGroupList);
                                setSelectedGroupList(0);
                                onClose();
                            }} colorScheme='red'>Delete</Button>
                        </ModalFooter>
                    </ModalContent>
                </Modal>
            </Stack>
        </Box>

    )
}
