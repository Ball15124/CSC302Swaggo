import {Box, Button, Editable, Textarea} from "@chakra-ui/react";

export const TodoListPage = () => {
    return (
        <Box>
            <Box>
                <Textarea placeholder={"Add a todo"} />
                <Button w={"100%"}>Add</Button>
            </Box>

            <Box p={10} bg={"gray.100"} gap={10} display={"flex"} flexDirection={"column"} rounded={"xl"} shadow={"lg"} >
                <Editable.Root defaultValue={"Todo 1"} >
                    <Editable.Preview />
                    <Editable.Input />
                </Editable.Root>
                <Button bg={"red"} >Delete</Button>
            </Box>


        </Box>
    );
};