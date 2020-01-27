package com.phodal.pholedge.book.model.command;

import lombok.Value;

import javax.validation.constraints.Size;

@Value
public class UpdateBookCommand {
    @Size(min = 5, message = "ISBN 长度大于 5")
    private String isbn;

    @Size(min = 1, message = "书名最小的长度大于 1")
    private String name;
}
