package com.phodal.pholedge.book.model.command;

import lombok.Value;

import javax.validation.constraints.NotNull;

@Value
public class CreateBookCommand {
    @NotNull(message = "ISBN 不能为空")
    private String isbn;

    @NotNull(message = "书名不能为空")
    private String name;
}
