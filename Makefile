# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: cpieri <cpieri@student.42.fr>              +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2018/03/15 11:20:25 by cpieri            #+#    #+#              #
#    Updated: 2020/05/20 11:19:46 by cpieri           ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

PACKAGE	=	expert-system

SRC_PATH=	src

GO		=	go

SRC_NAME= 	main.go		\
			parsing.go	\
			tools.go	\
			solver.go	\
			tree.go		\
			global.go

SRC		=	$(addprefix $(SRC_PATH)/,$(SRC_NAME))

NONE = \033[0m
RED = \033[31m
GREEN = \033[32m
YELLOW = \033[33m
BLUE = \033[34m
MAGENTA = \033[35m
CYAN = \033[36m

.PHONY:	all clean fclean re echo

all:
			@$(GO) build -o $(PACKAGE) $(SRC)
			@echo "$(GREEN)$(PACKAGE) ready!$(NONE)"

clean:
			@echo "$(YELLOW)Cleaning...$(NONE)"
			@/bin/rm -f $(PACKAGE)
			@echo "$(RED)$(PACKAGE) deleted !$(NONE)"

fclean:		clean

re:			fclean all
