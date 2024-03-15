import {SendPost} from "../../request/Api";
import React, {useEffect, useState} from "react";
import {
    Box,
    Button, Grid, GridItem, List,
    ListItem
} from "@chakra-ui/react";

export function ChannelsSubscriptionsTab() {
    let [channels, setChannels] = useState([])
    let [subscriptions, setSubscriptions] = useState([])
    let [currentChannelId, setCurrentChannelId] = useState(-1)

    useEffect(() => {
        SendPost('api/channels', {}).then(r => {
            setChannels(r.data)
            console.log(r.data)
        })
        SendPost('api/channels/subscriptions', {}).then(r => {
            setSubscriptions(r.data)
            console.log(r.data)
        })
    }, []);

    return (
        <Box>
            <Grid
                h='200px'
                templateRows='repeat(2, 1fr)'
                templateColumns='repeat(5, 1fr)'
                gap={4}
            >
                <GridItem rowSpan={2} colSpan={1}>
                        <List>
                            {
                                channels?.map((v) => <ListItem itemID={v.ID} key={v.ID}>
                                    <Button width="-webkit-fill-available" borderRadius={""} onClick={()=> {setCurrentChannelId(v.ID)}}>{v.Name}</Button>
                                </ListItem>)
                            }
                        </List>
                </GridItem>
                {subscriptions?.filter(v => v.ChannelsID === currentChannelId).map((v) => <GridItem colSpan={4} bg='lightgrey'>
                    <p>Type: {v.ChatType}</p>
                    <p>Name: {v.Username} {v.Title}</p>
                </GridItem>)}

            </Grid>
        </Box>

    )
}
