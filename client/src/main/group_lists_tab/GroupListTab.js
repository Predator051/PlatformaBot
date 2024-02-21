import {SendPost} from "../../request/Api";
import React, {useEffect, useState} from "react";
import {Box, Button, Input, InputGroup, InputRightElement, ListItem, OrderedList, Stack} from "@chakra-ui/react";

export function GroupListTab() {
    let [groupLists, setGroupLists] = useState([])
    let [newGroupList, setNewGroupList] = useState("")

    useEffect(() => {
        SendPost('api/group_lists', {}).then(r => {
            setGroupLists(r.data)
            console.log(r.data)
        })
    }, []);

    const createNewGroupList = (name) => {
        SendPost("api/new/group_lists", {name}).then(value => {
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
                    <Button colorScheme='blue' onClick={() => {createNewGroupList(newGroupList)}}>Create</Button>
                </InputGroup>
                <OrderedList>
                    {
                        groupLists.map((v) => <ListItem itemID={v.ID} key={v.ID}>{v.Name}</ListItem>)
                    }
                </OrderedList>
            </Stack>
        </Box>

    )
}
