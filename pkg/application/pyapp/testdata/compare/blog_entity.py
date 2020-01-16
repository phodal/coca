from ddd.shared.domain_model import DomainModel


class Blog(object):
    def __init__(self, id, title, content):
        self.id = id
        self.title = title
        self.content = content

    @classmethod
    def from_dict(cls, adict):
        blog = Blog(
            id=adict['id'],
            title=adict['title'],
            content=adict['content'],
        )

        return blog

    def to_dict(self):
        return {
            'id': self.id,
            'title': self.title,
            'content': self.content,
        }

    def __eq__(self, other):
        return self.to_dict() == other.to_dict()

DomainModel.register(Blog)
