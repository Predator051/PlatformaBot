import {SendPost} from "../../request/Api";
import React, {useEffect, useState} from "react";
import {MdDelete, MdMenu, MdOutlineSend} from "react-icons/md"
import {
    Box,
    Button,
    Divider,
    Flex, IconButton,
    Input,
    InputGroup,
    ListItem,
    Menu,
    MenuButton, MenuItem, MenuList,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent,
    ModalFooter,
    ModalHeader,
    ModalOverlay,
    OrderedList,
    Spacer,
    Stack, Textarea,
    useDisclosure
} from "@chakra-ui/react";

export function ChannelsTab() {
    let [channels, setChannels] = useState([])
    let [newChannel, setNewChannel] = useState("")
    const [sentMsgText, setSentMsgText] = useState("")
    const [selectedChannel, setSelectedChannel] = useState(0)
    const deleteModal = useDisclosure()
    const sendMsgModal = useDisclosure()

    useEffect(() => {
        SendPost('api/channels', {}).then(r => {
            setChannels(r.data)
            console.log(r.data)
        })
    }, []);

    const createNewChannel = (name) => {
        SendPost("api/new/channels", {name}).then(value => {
            SendPost('api/channels', {}).then(r => {
                setChannels(r.data ? r.data : [])
                console.log(r.data)
            })
        }, error => {})
    }

    const clickDeleteChannel = (id) => {
        if (id === 0) {return}

        console.log("Delete by id: ", id)
        SendPost("api/delete/channels", {id}).then(value => {
            SendPost('api/channels', {}).then(r => {
                setChannels(r.data)
                console.log(r.data)
            })
        }, error => {})
    }

    const clickSentMsgToChannel = (channels_id, msg) => {
        if (channels_id === 0 || msg === "") {return}

        console.log("Delete by id: ", channels_id);
        SendPost("api/channels/send/msg", {channels_id, msg});
    }

    return (
        <Box>
            <Stack spacing={3}>
                <InputGroup size='md'>
                    <Input
                        pr='4.5rem'
                        type={'text'}
                        placeholder='Enter new channel'
                        value={newChannel}
                        onChange={(event) => setNewChannel(event.target.value)}
                    />
                    <Button colorScheme='blue' onClick={() => {createNewChannel(newChannel); setNewChannel("")}}>Create</Button>
                </InputGroup>
                <OrderedList>
                    {
                        channels?.map((v) => <ListItem itemID={v.ID} key={v.ID}>
                            <Flex height={10}>
                                {v.Name}
                                <Spacer/>
                                <Menu>
                                    <MenuButton
                                        as={IconButton}
                                        aria-label='Options'
                                        icon={<MdMenu />}
                                        variant='outline'
                                    />
                                    <MenuList>
                                        <MenuItem icon={<MdDelete />} onClick={() => {
                                            setSelectedChannel(v.ID);
                                            deleteModal.onOpen();
                                        }}>
                                            Delete
                                        </MenuItem>
                                        <MenuItem icon={<MdOutlineSend />} onClick={() => {
                                            setSelectedChannel(v.ID);
                                            sendMsgModal.onOpen();
                                        }}>
                                            Send msg
                                        </MenuItem>
                                    </MenuList>
                                </Menu>
                            </Flex>
                            <Divider/>
                        </ListItem>)
                    }
                </OrderedList>
                <Modal onClose={deleteModal.onClose} isOpen={deleteModal.isOpen} isCentered>
                    <ModalOverlay />
                    <ModalContent>
                        <ModalHeader>Are you sure?</ModalHeader>
                        <ModalCloseButton />
                        <ModalBody>
                            All admins and connected groups will be deleted!
                        </ModalBody>
                        <ModalFooter>
                            <Button onClick={deleteModal.onClose} mr={3}>Close</Button>
                            <Button onClick={() => {
                                clickDeleteChannel(selectedChannel);
                                setSelectedChannel(0);
                                deleteModal.onClose();
                            }} colorScheme='red'>Delete</Button>
                        </ModalFooter>
                    </ModalContent>
                </Modal>
                <Modal onClose={sendMsgModal.onClose} isOpen={sendMsgModal.isOpen} isCentered>
                    <ModalOverlay />
                    <ModalContent>
                        <ModalHeader>Enter msg</ModalHeader>
                        <ModalCloseButton />
                        <ModalBody>
                            <Textarea placeholder='Here is a sample msg text' onChange={(e) => {
                                let inputValue = e.target.value
                                setSentMsgText(inputValue)
                            }} value={sentMsgText} />
                        </ModalBody>
                        <ModalFooter>
                            <Button onClick={sendMsgModal.onClose} mr={3}>Close</Button>
                            <Button onClick={() => {
                                clickSentMsgToChannel(selectedChannel, sentMsgText);
                                setSelectedChannel(0);
                                setSentMsgText("");
                                sendMsgModal.onClose();
                            }} colorScheme='green' isDisabled={sentMsgText === ""}>Send</Button>
                        </ModalFooter>
                    </ModalContent>
                </Modal>
            </Stack>
        </Box>

    )
}
