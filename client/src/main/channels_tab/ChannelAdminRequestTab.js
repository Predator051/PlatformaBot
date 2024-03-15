import {SendPost} from "../../request/Api";
import React, {useEffect, useState} from "react";
import {MdClose, MdDeck, MdDelete, MdDone} from "react-icons/md"
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

export function ChannelAdminRequestTab() {
    let [channelRequests, setChannelRequests] = useState([])
    let [channels, setChannels] = useState([])

    useEffect(() => {
        SendPost('api/channels/admin/requests', {}).then(r => {
            setChannelRequests(r.data.requests)
            setChannels(r.data.channels)
            console.log(r.data)
        })
    }, []);

    const clickAcceptRequest = (channelRequest) => {
        console.log(channelRequest)
        SendPost("api/channels/admin/accept",
            {
                channels_id: channelRequest.ChannelsID,
                chat_id: channelRequest.ChatID
            }).then(value => {
            SendPost('api/channels/admin/requests', {}).then(r => {
                setChannelRequests(r.data.requests)
                setChannels(r.data.channels)
                console.log(r.data)
            })
        }, error => {
        })
    }

    return (
        <Box>
            <Stack spacing={3}>
                <OrderedList>
                    {
                        channelRequests?.map((v) => <ListItem itemID={v.ID} key={v.ID}>
                            <Flex height={10}>
                                User {' '}
                                [{v.Username} {' '}
                                {v.FirstName} {' '}
                                {v.SecondName} {' '}]
                                requests for administration of {' '}
                                {channels?.find(gl => gl.ID == v.ChannelsID)?.Name}
                                <Spacer/>
                                <Button colorScheme='green' leftIcon={<MdDone/>} mr={1} onClick={()=>clickAcceptRequest(v)}>
                                    Accept
                                </Button>
                                <Button colorScheme='red' leftIcon={<MdClose/>}>
                                    Decline
                                </Button>
                            </Flex>
                            <Divider/>
                        </ListItem>)
                    }
                </OrderedList>
            </Stack>
        </Box>

    )
}
