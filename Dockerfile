FROM scratch
COPY jcconv /
ENTRYPOINT ["/jcconv"]
CMD ["web"]
